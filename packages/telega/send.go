package telega

import (
	"barcode/packages/must"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Telega) Send(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	must.Panic2(c.bot.Send(msg))
}
