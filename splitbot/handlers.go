package splitbot

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/almiskov/text-split-bot/splitter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *splitbot) handleCommand(msg *tgbotapi.Message) {
	switch msg.Text {
	case "/start":
		replyText := fmt.Sprintf(
			"Привет, %s, нажми \"%s\" и отправляй текст 🙂",
			msg.From.FirstName,
			startKbText,
		)
		reply := tgbotapi.NewMessage(msg.Chat.ID, replyText)
		reply.ReplyMarkup = startKeyboard
		b.bot.Send(reply)

		b.store.Clear(msg.From.ID)
	}
}

func (b *splitbot) handleStartKbText(msg *tgbotapi.Message) {
	replyText := fmt.Sprintf(
		"%s, отправь столько текста, сколько нужно (можно в нескольких сообщениях), после чего нажми \"%s\"",
		msg.From.FirstName,
		processKbText,
	)

	b.store.Init(msg.From.ID)

	reply := tgbotapi.NewMessage(msg.Chat.ID, replyText)
	reply.ReplyMarkup = processKeyboard
	b.bot.Send(reply)
}

func (b *splitbot) handleProcessKbText(msg *tgbotapi.Message) {
	replyText, ok := b.store.Get(msg.From.ID)
	if !ok || replyText == "" {
		replyText = fmt.Sprintf("%s, отправь хотя бы одно сообщение ☝", msg.From.FirstName)
		reply := tgbotapi.NewMessage(msg.Chat.ID, replyText)
		b.bot.Send(reply)
		return
	}

	spl := splitter.New()
	pages := spl.Split(replyText)

	sendBorder := func(d BorderDirection) {
		border := tgbotapi.NewMessage(msg.Chat.ID, d.String())
		b.bot.Send(border)
	}

	sendBorder(Down)

	lengths := make([]string, len(pages))

	for i, page := range pages {
		lengths[i] = strconv.Itoa(utf8.RuneCountInString(page))

		reply := tgbotapi.NewMessage(msg.Chat.ID, page)
		b.bot.Send(reply)
	}

	sendBorder(Up)

	finalMsg := tgbotapi.NewMessage(msg.Chat.ID, "Страницы получились такой длины: "+strings.Join(lengths, ", "))
	finalMsg.ReplyMarkup = startKeyboard
	b.bot.Send(finalMsg)

	b.store.Clear(msg.From.ID)
}

func (b *splitbot) handleRestartKbText(msg *tgbotapi.Message) {
	b.store.Init(msg.From.ID)

	replyText := fmt.Sprintf("Начнём заново, %s 🙂", msg.From.FirstName)

	reply := tgbotapi.NewMessage(msg.Chat.ID, replyText)
	reply.ReplyMarkup = processKeyboard
	b.bot.Send(reply)
}

func (b *splitbot) handleDefaultMessage(msg *tgbotapi.Message) {
	_, ok := b.store.Get(msg.From.ID)
	if ok {
		b.store.Add(msg.From.ID, msg.Text)
		return
	}

	replyText := fmt.Sprintf("Это сообщение проигнорировано 😟 Нажми \"%s\"", startKbText)

	reply := tgbotapi.NewMessage(msg.Chat.ID, replyText)
	b.bot.Send(reply)
}
