package logic

import "barcode/packages/telega"

func StatusCmdReply(bot *telega.Telega, chatId int64) {
	bot.Send(chatId, "I am working")
}
