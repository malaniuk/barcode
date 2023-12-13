package telega

import (
	"barcode/packages/must"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Telega) GetFileLink(id string) string {
	opts := tgbotapi.FileConfig{FileID: id}
	file := must.Panic2(c.bot.GetFile(opts))

	return file.Link(c.apiKey)
}
