package pluginGroupManager

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"strconv"
)

var (
	plugin *leafBot.Plugin
)

func init() {
	Init()
	plugin = leafBot.NewPlugin("q群管理")
}
func Init() {
	plugin.OnRegex(`^升为管理.*?qq=(\d+)`).
		SetPluginName("群管系统-设置管理").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					bot.Send(event, message.Text("发生未知错误"+err.Error()))
					return
				}
				bot.SetGroupAdmin(event.GroupId, ID, true)
				nickName := bot.GetGroupMemberInfo(event.GroupId, ID, true).NickName
				bot.Send(event, message.Text(nickName+"升为了管理！"))
			})

	plugin.OnRegex(`^取消管理.*?qq=(\d+)`).
		SetPluginName("群管系统-取消管理").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					bot.Send(event, message.Text("发生未知错误"+err.Error()))
					return
				}
				bot.SetGroupAdmin(event.GroupId, ID, false)
				nickName := bot.GetGroupMemberInfo(event.GroupId, ID, true).NickName
				bot.Send(event, message.Text(nickName+"失去了管理员的资格！"))
			},
		)

	plugin.OnRegex(`^踢出群聊.*?qq=(\d+)`).
		SetPluginName("群管系统-踢出群聊").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					bot.Send(event, message.Text("发生未知错误"+err.Error()))
					return
				}
				bot.SetGroupKick(event.GroupId, ID, false)
				nickName := bot.GetGroupMemberInfo(event.GroupId, ID, true).NickName
				bot.Send(event, message.Text(nickName+"被移除了群聊！"))
			},
		)

	plugin.OnRegex(`^退出群聊.*?(\d+)`).
		SetPluginName("群管系统-退出群聊").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					bot.Send(event, message.Text("发生未知错误"+err.Error()))
					return
				}
				bot.SetGroupLeave(ID, true)
				nickName := bot.GetGroupMemberInfo(event.GroupId, ID, true).NickName
				bot.Send(event, message.Text(nickName+"被移除了群聊！"))
			},
		)

	plugin.OnCommand(`开启全员禁言`).
		SetPluginName("群管系统-全体禁言").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {

				bot.SetGroupWholeBan(event.GroupId, true)
				bot.Send(event, message.Text("全员开始自闭"))
			},
		)

	plugin.OnCommand(`解除全员禁言`).
		SetPluginName("群管系统-关闭全员群聊").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {

				bot.SetGroupWholeBan(event.GroupId, false)
				bot.Send(event, message.Text("全员自闭结束"))
			},
		)

}
