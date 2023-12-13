package logic

import (
	"barcode/config"
	"barcode/packages/telega"
	"fmt"
	"github.com/samber/lo"
)

func IsAllowedIds(conf *config.Config, bot *telega.Telega, chatId int64) bool {
	if lo.Contains(conf.TgAllowedIds, chatId) {
		return true
	}

	msgText := fmt.Sprintf("You are not allowed to use this bot\nYour chatId is %v", chatId)
	bot.Send(chatId, msgText)

	return false
}
