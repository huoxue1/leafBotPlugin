package plugin_gif

import (
	"fmt"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"github.com/huoxue1/leafBotPlugin/plugin_gif/gif"
)

var (
	m = map[string]func(string) string{
		"摸": gif.Mo,
		"搓": gif.Cuo,
		"冲": gif.Chong,
		"拍": gif.Pai,
	}
)

func MoInit() {

	leafBot.OnMessage("group").SetWeight(10).SetPluginName("gif").AddRule(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) bool {
		for s, _ := range m {
			if event.Message[0].Type == "text" && event.Message[0].Data["text"] == s {
				state.Data["type"] = event.Message[0].Data["text"]
				for _, segment := range event.Message {
					if segment.Type == "at" {
						state.Data["data"] = segment.Data["qq"]
						return true
					}
				}
			}
		}
		return false
	}).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		f := m[state.Data["type"].(string)]
		bot.Send(event, message.Image(f(fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%v&s=100", state.Data["data"]))))
	})

}
