package logic

import (
	"barcode/packages/must"
	"barcode/packages/telega"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/samber/lo"
	"image/jpeg"
	"log"
	"net/http"
)

func BarCodeUploaded(bot *telega.Telega, chatId int64, photos []tgbotapi.PhotoSize) string {
	lastPhoto := must.Panic2(lo.Last(photos))
	fileLink := bot.GetFileLink(lastPhoto.FileID)

	log.Println(fileLink)

	fileResp := must.Panic2(http.Get(fileLink))
	defer func() {
		must.Panic(fileResp.Body.Close())
	}()

	img := must.Panic2(jpeg.Decode(fileResp.Body))
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	reader := oned.NewCode39Reader()

	barCode, err := reader.DecodeWithoutHints(bmp)
	if err != nil {
		if err.Error() == "NotFoundException" {
			bot.Send(chatId, "Bar code not found")
			return ""
		}

		log.Print(err)
	}

	msg := fmt.Sprintf("Bar code is %v", barCode)
	bot.Send(chatId, msg)
	bot.Send(chatId, "Send images that needs to be renamed and packed, and run /pack command after")

	return barCode.GetText()
}
