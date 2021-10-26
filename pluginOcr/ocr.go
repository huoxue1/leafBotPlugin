package pluginOcr

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	go ocr()
}

// Ocr
/**
 * @Description:
 * example
 */
func ocr() {
	plugin := leafBot.NewPlugin("图片ocr")
	plugin.SetHelp(map[string]string{
		"ocr": "图片ocr",
	})
	plugin.OnCommand("ocr").SetPluginName("图片ocr").SetBlock(false).SetWeight(10).AddRule(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) bool {
		for _, mess := range event.Message {
			if mess.Type == "image" {
				return true
			}
		}
		return false
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		images := event.GetImages()
		if len(images) < 1 {
			return
		}
		ocrImage := bot.(leafBot.OneBotApi).OcrImage(images[0].Data["file"])
		mess := "识别结果:\n识别语言:" + ocrImage.Get("language").String()
		for _, text := range ocrImage.Get("texts").Array() {
			mess += "\n" + text.Get("text").String()
		}

		event.Send(message.ReplyWithMessage(int64(event.MessageID), message.Text(mess)))
	})
}
