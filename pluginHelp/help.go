package pluginHelp

import (
	"encoding/base64"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"github.com/huoxue1/leafBot/utils"
)

func init() {
	leafBot.OnCommand("/help").
		AddAllies("帮助").
		AddRule(leafBot.OnlyToMe).
		SetPluginName("帮助").
		SetBlock(false).
		SetCD("default", 0).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				event.Send(message.Text("downloading image ......"))
				screen, err := utils.GetPWScreen("https://huoxue1.github.io/leafBot/Features", "")
				if err != nil {
					event.Send(message.Text("获取帮助文档失败" + err.Error()))
					return
				}
				event.Send(message.Image("base64://" + base64.StdEncoding.EncodeToString(screen)))
			})
}
