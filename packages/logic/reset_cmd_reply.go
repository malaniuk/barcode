package logic

import "barcode/packages/telega"

func ResetCmdReply(bot *telega.Telega, chatId int64) {
	msgText := "Reset successfully"
	bot.Send(chatId, msgText)
}
