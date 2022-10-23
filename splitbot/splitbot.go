package splitbot

import (
	"log"

	"github.com/almiskov/text-split-bot/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type splitbot struct {
	bot   *tgbotapi.BotAPI
	store *store.Store
}

func New(token string, debug bool) (*splitbot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &splitbot{
		bot:   bot,
		store: store.New(),
	}, nil
}

func (b *splitbot) Run() {
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	for upd := range b.bot.GetUpdatesChan(ucfg) {
		b.handleUpdate(upd)
	}
}

func (b *splitbot) handleUpdate(upd tgbotapi.Update) {
	switch true {
	case upd.Message != nil:
		handler := b.handleMessage
		if upd.Message.IsCommand() {
			handler = b.handleCommand
		}
		handler(upd.Message)
	default:
		log.Println("unknown upd")
	}
}

func (b *splitbot) handleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	case startKbText:
		b.handleStartKbText(msg)
	case processKbText:
		b.handleProcessKbText(msg)
	case restartKbText:
		b.handleRestartKbText(msg)
	default:
		b.handleDefaultMessage(msg)
	}
}
