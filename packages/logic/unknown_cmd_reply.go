package logic

import "barcode/packages/telega"

func UnknownCmdReply(bot *telega.Telega, chatId int64) {
	bot.Send(chatId, "Unknown command")
}
