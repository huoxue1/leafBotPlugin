package plugin_fakeMsg

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"strconv"
	"strings"
)

func init() {
	FakeMsg()
}

func FakeMsg() {
	plugin := leafBot.NewPlugin("假消息")
	plugin.SetHelp(map[string]string{"fakeMsg": "获取一条合并转发的假消息"})
	plugin.OnCommand("fakeMsg").
		SetBlock(false).
		SetPluginName("假消息").
		AddAllies("假消息").
		SetCD("default", 3).
		SetWeight(10).
		AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
			bot.Send(event, message.Text("请发送需要发送的假消息,发送`《《`停止录入信息：\neg:\n昵称 qq号 消息"))
			messages := message.Message{}
			for true {
				event1, err := bot.GetOneEvent(func(event1 leafBot.Event, bot1 *leafBot.Bot, state1 *leafBot.State) bool {
					if event1.GroupId == event.GroupId && event1.UserId == event.UserId {
						return true
					}
					return false
				})
				if err != nil || event.Message.ExtractPlainText() == "《《" {
					break
				}
				text := event1.GetPlainText()
				datas := strings.Split(text, " ")
				if len(datas) < 3 {
					bot.Send(event, message.Text("错误的信息录入"))
					break
				}
				id, err := strconv.ParseInt(datas[1], 10, 60)
				if err != nil {
					break
				}
				messages = append(messages, message.CustomNode(datas[0], id, datas[2]))

			}
			if len(messages) < 1 {
				return
			}

			bot.SendGroupForwardMsg(event.GroupId, messages)

		})
}
