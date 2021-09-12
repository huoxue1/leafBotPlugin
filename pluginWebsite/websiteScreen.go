package pluginWebsite

import (
	"encoding/base64"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"github.com/huoxue1/leafBot/utils"
)

func init() {
	WebSiteScreenInit()
}

func WebSiteScreenInit() {
	plugin := leafBot.NewPlugin("网页长截图")
	plugin.SetHelp(map[string]string{">website": "对指定网页截图"})
	plugin.OnCommand(">website").AddAllies("网页截图").SetPluginName("网页长截图").SetCD("default", 0).SetBlock(false).SetWeight(10).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		if len(state.Args) < 1 {
			bot.Send(event, message.Text("参数不足，详情参考帮助菜单"))
			return
		}
		data, err := utils.GetPWScreen(state.Args[0], "")
		if err != nil {
			bot.Send(event, message.Text("获取截图错误"+err.Error()))
			return
		}
		bot.Send(event, message.Image("base64://"+base64.StdEncoding.EncodeToString(data)))
	})
}
