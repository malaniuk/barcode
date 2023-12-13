package telega

import (
	"barcode/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Telega struct {
	apiKey string

	bot *tgbotapi.BotAPI
}

func NewTelega(conf *config.Config) *Telega {
	bot, err := tgbotapi.NewBotAPI(conf.TgApiKey)
	if err != nil {
		log.Fatalf("Bot init error %v", err)
	}

	bot.Debug = true

	log.Printf("Bot name %s", bot.Self.UserName)

	return &Telega{
		bot:    bot,
		apiKey: conf.TgApiKey,
	}
}
