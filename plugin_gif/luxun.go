package plugin_gif

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"github.com/guonaihong/gout"
	"github.com/huoxue1/gg"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	log "github.com/sirupsen/logrus"
	"image"
	"strings"
)

var font []byte

func init() {
	LuXun()

	err := gout.GET("https://specialblog.link/img/202109091139659.ttf").BindBody(&font).Do()
	if err == nil {
		log.Infoln("加载字体文件成功")
	}
}

func LuXun() {
	leafBot.OnStartWith("鲁迅说").SetPluginName("鲁迅说").SetWeight(10).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		text := event.GetPlainText()
		data := strings.TrimLeft(text, "鲁迅说")
		if len(data) == 0 {
			bot.Send(event, message.Text("你想让鲁迅说点什么呢？"))
			event1, err := bot.GetOneEvent(func(event1 leafBot.Event, bot1 *leafBot.Bot, state1 *leafBot.State) bool {
				if event1.UserId == event.UserId && event1.GroupId == event.GroupId {
					return true
				}
				return false
			})
			if err != nil {
				return
			}
			data = event1.GetPlainText()
		}

		img, err := getImage(data)
		if err != nil {
			bot.Send(event, message.Text("鲁迅说出错了"+err.Error()))
			return
		}
		bot.Send(event, message.Image("base64://"+base64.StdEncoding.EncodeToString(img)))
	})
}

func getImage(text string) ([]byte, error) {
	var result []byte
	buffer := bytes.NewBuffer(result)
	var data []byte
	err := gout.GET("https://specialblog.link/img/202109090936718.jpeg").BindBody(&data).Do()
	if err != nil {
		return nil, err
	}
	decode, s, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	log.Infoln(s)
	context := gg.NewContextForImage(decode)
	err = context.LoadFontFromBytes(font, 30)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	context.SetHexColor("FFFFFF")
	context.DrawString("——鲁迅", 320, 440)
	log.Infoln(len(text))
	context.DrawString(text, 240-float64(len(text)*5), 370)
	err = context.EncodePNG(buffer)
	if err != nil {
		return nil, err
	}

	return result, err

}
