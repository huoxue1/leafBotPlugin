package pluginQrCode

import (
	"fmt" //nolint:gci
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	UseCreateQrCode()
}

// UseCreateQrCode 生成二维码的插件
func UseCreateQrCode() {

	plugin := leafBot.NewPlugin("二维码生成")
	plugin.SetHelp(map[string]string{
		"/createQrcode": "生成二维码",
	})
	plugin.OnCommand("/createQrcode").
		AddAllies("生产二维码").
		SetWeight(10).
		SetPluginName("二维码生成").
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				switch len(state.Args) {
				case 0:
					{
						bot.Send(event, "参数不足")
					}
				case 1:
					{
						bot.Send(event, message.Image(fmt.Sprintf("https://api.isoyu.com/qr/?m=0&e=L&p=15&url=%s", state.Args[0])).Add("c", 3).Add("cache", 0))
					}
				case 2:
					{
						bot.Send(event, message.Image(fmt.Sprintf("https://api.isoyu.com/qr/?m=%v&e=L&p=15&url=%s", state.Args[1], state.Args[0])).Add("cache", 0).Add("c", 3))
					}
				}
			})

}
