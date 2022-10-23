package splitbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	startKbText   string = "Начать 🚩"
	processKbText string = "Обработать 💫"
	restartKbText string = "Начать заново 🏳"
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(startKbText),
	),
)

var processKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(processKbText),
		tgbotapi.NewKeyboardButton(restartKbText),
	),
)
