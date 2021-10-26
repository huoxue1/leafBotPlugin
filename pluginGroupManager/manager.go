package pluginGroupManager

import (
	"strconv"

	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
)

var (
	plugin *leafBot.Plugin
)

func init() {
	go Init()
	plugin = leafBot.NewPlugin("q群管理")
}
func Init() {
	plugin.OnRegex(`^升为管理.*?qq=(\d+)`).
		SetPluginName("群管系统-设置管理").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					event.Send(message.Text("发生未知错误" + err.Error()))
					return
				}
				bot.(leafBot.OneBotApi).SetGroupAdmin(event.GroupId, ID, true)
				nickName := bot.(leafBot.OneBotApi).GetGroupMemberInfo(event.GroupId, ID, true).Get("nick_name").String()
				event.Send(message.Text(nickName + "升为了管理！"))
			})

	plugin.OnRegex(`^取消管理.*?qq=(\d+)`).
		SetPluginName("群管系统-取消管理").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					event.Send(message.Text("发生未知错误" + err.Error()))
					return
				}
				bot.(leafBot.OneBotApi).SetGroupAdmin(event.GroupId, ID, false)
				nickName := bot.(leafBot.OneBotApi).GetGroupMemberInfo(event.GroupId, ID, true).Get("nick_name").String()
				event.Send(message.Text(nickName + "失去了管理员的资格！"))
			},
		)

	plugin.OnRegex(`^踢出群聊.*?qq=(\d+)`).
		SetPluginName("群管系统-踢出群聊").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					event.Send(message.Text("发生未知错误" + err.Error()))
					return
				}
				bot.(leafBot.OneBotApi).SetGroupKick(event.GroupId, ID, false)
				nickName := bot.(leafBot.OneBotApi).GetGroupMemberInfo(event.GroupId, ID, true).Get("nick_name").String()
				event.Send(message.Text(nickName + "被移除了群聊！"))
			},
		)

	plugin.OnRegex(`^退出群聊.*?(\d+)`).
		SetPluginName("群管系统-退出群聊").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					event.Send(message.Text("发生未知错误" + err.Error()))
					return
				}
				bot.(leafBot.OneBotApi).SetGroupLeave(ID, true)
				nickName := bot.(leafBot.OneBotApi).GetGroupMemberInfo(event.GroupId, ID, true).Get("nick_name").String()
				event.Send(message.Text(nickName + "被移除了群聊！"))
			},
		)

	plugin.OnCommand(`开启全员禁言`).
		SetPluginName("群管系统-全体禁言").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				bot.(leafBot.OneBotApi).SetGroupWholeBan(event.GroupId, true)
				event.Send(message.Text("全员开始自闭"))
			},
		)

	plugin.OnCommand(`解除全员禁言`).
		SetPluginName("群管系统-关闭全员群聊").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
				bot.(leafBot.OneBotApi).SetGroupWholeBan(event.GroupId, false)
				event.Send(message.Text("全员自闭结束"))
			},
		)
}
