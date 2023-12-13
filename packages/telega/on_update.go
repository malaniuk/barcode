package telega

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Telega) OnUpdates(callBack func(message *tgbotapi.Message)) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := c.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() && len(update.Message.Photo) == 0 {
			continue
		}

		callBack(update.Message)
	}
}
