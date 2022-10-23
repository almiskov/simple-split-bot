package splitter

import (
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/almiskov/text-split-bot/internal"
)

const (
	long  = "examples/long.txt"
	short = "examples/short.txt"
)

func TestSplitV2(t *testing.T) {

	text := string(internal.Must(os.ReadFile(long)))

	spl := New(WithLeadingSize(500), WithRegularSize(300))

	sz := spl.leadingSize
	paragraphs := spl.Split(text)
	for i, p := range paragraphs {
		if i > 0 {
			sz = spl.regularSize
		}

		length := utf8.RuneCountInString(p)
		if length > sz {
			t.Errorf("len of page %d (%d) > allowed for this page (%d)", i, length, sz)
		}
	}

	f := internal.Must(os.Create("examples/result.txt"))
	defer f.Close()

	for _, p := range paragraphs {
		io.WriteString(f, p)
		io.WriteString(f, "\n")
		io.WriteString(f, "⤴"+strconv.Itoa(utf8.RuneCountInString(p)))
		io.WriteString(f, "\n")
	}
}

func TestInstaSecretSpace(t *testing.T) {
	rd := strings.NewReader("Этот невидимый пробел можно скопировать отсюда: «⠀⠀⠀⠀». (Их несколько штук между кавычками).")

	_ = rd

	t.Fail()
	for _, r := range " ⠀" {
		t.Log(r)
	}
}

func TestBrakeTextReg(t *testing.T) {

	text := string(internal.Must(os.ReadFile(long)))

	spl := New()

	parts := spl.breakTextReg(text)

	t.Fail()

	for _, p := range parts {
		t.Log(p)
	}
}
