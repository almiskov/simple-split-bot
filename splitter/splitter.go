package splitter

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	instaNL = "\n⠀\n"
)

type splitter struct {
	leadingSize, regularSize int
	pageTmpl, nextTmpl       string
	nextTmplLen              int
	leadingPageAddition      string
}

func New(opts ...splitterOpt) *splitter {
	const (
		defaultLeading             = 2200
		defaultRegular             = 1000
		defaultPageTmpl            = "❗%d❗"
		defaultNextTmpl            = "👇👇👇"
		defaultLeadingPageAddition = ""
	)

	s := &splitter{
		leadingSize:         defaultLeading,
		regularSize:         defaultRegular,
		pageTmpl:            defaultPageTmpl,
		nextTmpl:            defaultNextTmpl,
		leadingPageAddition: defaultLeadingPageAddition,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.nextTmplLen = utf8.RuneCountInString(s.nextTmpl)

	return s
}

/*
breakText разбивает исходный текст на части по ключевым символам (см. switch-case).
Обрезает все пробелы в каждом куске текста. Перенос строки оставляет в том же куске,
после которого он находится
*/
func (s *splitter) breakText(text string) []string {
	var (
		parts      = make([]string, 0)
		start, end = 0, 0
	)

	for i, r := range text {
		switch r {
		case '.', ',', '!', '?', ';', ':', '—':
			end = i + utf8.RuneLen(r)
			parts = append(parts, strings.TrimLeft(text[start:end], " "))
			start = end
		}
	}

	return append(parts, strings.TrimLeft(text[start:], " "))
}

var brakeReg = regexp.MustCompile(`[.|,|!|?|;|:|—]+\s`)

/*
breakTextReg разбивает исходный текст на части по регулярке, по кон
*/
func (s *splitter) breakTextReg(text string) []string {
	idxs := brakeReg.FindAllStringIndex(text, -1)

	parts := make([]string, 0)
	start, end := 0, 0

	for _, p := range idxs {
		end = p[1]
		parts = append(parts, strings.Trim(text[start:end], " "))
		start = end
	}

	return append(parts, strings.Trim(text[start:], " "))
}

func (s *splitter) aggregatePartsToPages(parts []string) []string {
	var (
		pages         = make([]string, 0)
		curPageLength = utf8.RuneCountInString(s.leadingPageAddition + instaNL)
	)

	// собираем главную страницу
	for i, part := range parts {
		curPageLength += utf8.RuneCountInString(part)
		if curPageLength+i+s.nextTmplLen > s.leadingSize {
			pages = append(pages, strings.Join(parts[:i], " ")+s.leadingPageAddition+s.nextTmpl)
			parts = parts[i:]
			curPageLength = 0
			break
		}
	}

	var (
		pageTitle    string
		pageTitleLen int
	)

	// собираем остальные страницы
	for i := 0; i < len(parts); i++ {
		if i == 0 {
			pageTitle = fmt.Sprintf(s.pageTmpl, len(pages))
			pageTitleLen = utf8.RuneCountInString(pageTitle)
		}

		curPageLength += utf8.RuneCountInString(parts[i])

		if pageTitleLen+curPageLength+i+s.nextTmplLen > s.regularSize {
			pages = append(pages, pageTitle+strings.Join(parts[:i], " ")+s.nextTmpl)

			parts = parts[i:]
			curPageLength = 0
			i = -1 // дальше i++ превратит его в 0
		}
	}

	// добавляем последнюю страницу и возвращаем
	return append(pages, pageTitle+strings.Join(parts, " "))
}

func (s *splitter) replaceNewLines(text string) string {
	return strings.NewReplacer("\n\n", instaNL, "\r\n\r\n", instaNL).Replace(text)
}

/*
Split разбивает текст на страницы, размером не более leadingSize и regularSize,
при необходимости добавляя форматированный номер страницы (начиная с первой)
и указатель, что следующая страница существует, если она существует
*/
func (s *splitter) Split(text string) []string {
	text = s.replaceNewLines(text)

	if utf8.RuneCountInString(text)+utf8.RuneCountInString(s.leadingPageAddition) <= s.leadingSize {
		return []string{text + s.leadingPageAddition}
	}

	var (
		parts      = s.breakTextReg(text)
		paragraphs = s.aggregatePartsToPages(parts)
	)

	return paragraphs
}
