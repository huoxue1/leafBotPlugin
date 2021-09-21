package plugin_reply

import (
	"github.com/huoxue1/leafBot"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"os"
)

func init() {
	if InitFile() {
		InitReply()
	}
}

var data []byte

func InitFile() bool {
	file, err := os.OpenFile("/config/reply.json", os.O_RDWR, 0666)
	if err != nil {
		log.Errorln("未发现闲聊词库")
		return false
	}
	data, _ = io.ReadAll(file)
	return true
}

func InitReply() {
	plugin := leafBot.NewPlugin("reply")
	plugin.SetHelp(map[string]string{"根据词库生成闲聊": ""})
	plugin.OnMessage("group").SetPluginName("对话").AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		text := event.GetPlainText()
		mess := gjson.GetBytes(data, text)
		if mess.Exists() {
			event.Send(text)
		}
	})
}
