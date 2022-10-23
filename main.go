package main

import (
	"log"

	"github.com/almiskov/text-split-bot/splitbot"
)

func main() {
	var cfg config

	if err := cfg.Parse("config.json"); err != nil {
		log.Fatal(err)
	}

	bot, err := splitbot.New(cfg.Token, cfg.Debug)
	if err != nil {
		log.Fatal(err)
	}

	bot.Run()
}
