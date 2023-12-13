package main

import (
	"barcode/config"
	"barcode/packages/logic"
	"barcode/packages/telega"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	conf := config.NewConfig()
	bot := telega.NewTelega(conf)

	session := map[int64]string{}

	bot.OnUpdates(func(message *tgbotapi.Message) {
		chatId := message.Chat.ID
		//
		//if !logic.IsAllowedIds(conf, bot, chatId) {
		//	return
		//}

		if message.IsCommand() {
			switch message.Command() {
			case "start":
				logic.StartCmdReply(bot, chatId)
			case "status":
				logic.StatusCmdReply(bot, chatId)
			case "reset":
				delete(session, chatId)
				logic.ResetCmdReply(bot, chatId)
				logic.StartCmdReply(bot, chatId)
			default:
				logic.UnknownCmdReply(bot, chatId)
			}

			return
		}

		if len(message.Photo) > 0 {
			barCode := logic.BarCodeUploaded(bot, chatId, message.Photo)
			if len(barCode) == 0 {
				return
			}

			session[chatId] = barCode
		}

		log.Print(message)
	})
}
