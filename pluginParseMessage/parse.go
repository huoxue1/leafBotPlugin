package pluginParseMessage

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	InitParse()
}

//InitParse
/**
 * @Description:
 * example
 */
func InitParse() {
	plugin := leafBot.NewPlugin("消息解析")
	plugin.SetHelp(map[string]string{
		"decode": "解析消息，还原为cq码",
		"run":    "运行cq码",
	})
	plugin.OnCommand("decode").SetWeight(10).SetPluginName("消息解析").SetBlock(false).SetCD("default", 5).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		bot.Send(event, message.Text("请在一分钟内发出需要解析的消息"))
		oneEvent, err := bot.GetOneEvent(func(event1 leafBot.Event, bot *leafBot.Bot, state *leafBot.State) bool {
			if event1.GroupId == event.GroupId && event1.UserId == event.UserId {
				return true
			}
			return false
		})
		if err != nil {
			bot.Send(event, message.Text("等待超时"))
			return
		}
		bot.Send(event, message.Text(oneEvent.Message.CQString()))
	})

	plugin.OnCommand("run").SetWeight(10).SetPluginName("cq码运行").SetBlock(false).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		if len(state.Args) < 1 {
			return
		}
		bot.Send(event, message.ParseMessageFromString(state.Args[0]))
	})
}
