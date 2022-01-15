package group_manager

import (
	"fmt"
	"strconv"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
)

func init() {
	manager()
}

func manager() {
	plugin := leafbot.NewPlugin("group-manager")
	// 升为管理
	plugin.OnStart("升为管理", leafbot.Option{
		Weight: 1,
		Block:  true,
		Rules:  []leafbot.Rule{},
	}).Handle(func(ctx *leafbot.Context) {
		for _, v := range ctx.Event.Message {
			if v.Type == "at" {
				qq, _ := strconv.Atoi(v.Data["qq"])
				ctx.Bot.(leafbot.OneBotAPI).SetGroupAdmin(ctx.Event.GroupId, qq, true)
				ctx.Send(message.Text(fmt.Sprintf("%v已经升为管理员", qq)))
			}
		}
	})

	// 取消管理
	plugin.OnStart("取消管理", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		for _, v := range ctx.Event.Message {
			if v.Type == "at" {
				qq, _ := strconv.Atoi(v.Data["qq"])
				ctx.Bot.(leafbot.OneBotAPI).SetGroupAdmin(ctx.Event.GroupId, qq, false)
				ctx.Send(message.Text(fmt.Sprintf("%v已被取消管理员", qq)))
			}
		}
	})
	// 踢出群聊
	plugin.OnStart("踢出群聊", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		for _, v := range ctx.Event.Message {
			if v.Type == "at" {
				qq, _ := strconv.Atoi(v.Data["qq"])
				ctx.Bot.(leafbot.OneBotAPI).SetGroupKick(ctx.Event.GroupId, qq, false)
				ctx.Send(message.Text(fmt.Sprintf("%v已被踢出群聊", qq)))
			}
		}
	})

	// 踢出群聊
	plugin.OnRegex("^退出群聊(/d+)", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		id, _ := strconv.Atoi(ctx.State.RegexResult[1])
		ctx.SetGroupLeave(id, true)
		ctx.Send(message.Text("已退出群聊：%d", id))
	})
	// 开启全员禁言
	plugin.OnFullMatchGroup("开启全员禁言", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		ctx.SetGroupWholeBan(ctx.Event.GroupId, true)
	})

	// 关闭全员禁言
	plugin.OnFullMatchGroup("解除全员禁言", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		ctx.SetGroupWholeBan(ctx.Event.GroupId, false)
	})

	// 自闭
	plugin.OnRegex("[我要自闭|禅定](\\d+)(小时|天|分钟|day|min|h|)", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		duration, _ := strconv.Atoi(ctx.State.Args[1])
		switch ctx.State.Args[2] {
		case "天", "day":
			duration = duration * 60 * 60 * 24
		case "分钟", "min", "m":
			duration = duration * 60
		case "h", "小时":
			duration = duration * 60 * 60
		default:
			duration = duration * 60
		}
		ctx.SetGroupBan(ctx.Event.GroupId, ctx.Event.UserId, duration)
		ctx.Send(append(message.Message{}, message.Text("先去休息吧"), message.At(int64(ctx.Event.UserId))))
	})

	plugin.OnRegex("^申请头衔(.*?)", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		ctx.SetGroupSpecialTitle(ctx.Event.GroupId, ctx.Event.UserId, ctx.State.RegexResult[1], -1)
		ctx.Send(append(message.Message{}, message.Text("申请成功"), message.At(int64(ctx.Event.UserId))))
	})

}
