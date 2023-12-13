package logic

import "barcode/packages/telega"

func StartCmdReply(bot *telega.Telega, chatId int64) {
	msgText := "Upload image with a bar code to detect the number"
	bot.Send(chatId, msgText)
}
