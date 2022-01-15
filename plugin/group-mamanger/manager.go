package group_manager

import (
	"fmt"
	"strconv"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
	log "github.com/sirupsen/logrus"
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
		Rules:  []leafbot.Rule{UserSuperUser},
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
		Rules:  []leafbot.Rule{UserSuperUser},
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
		Rules:  []leafbot.Rule{UserAdminPermission},
	}).Handle(func(ctx *leafbot.Context) {
		for _, v := range ctx.Event.Message {
			if v.Type == "at" {
				qq, _ := strconv.Atoi(v.Data["qq"])
				ctx.Bot.(leafbot.OneBotAPI).SetGroupKick(ctx.Event.GroupId, qq, false)
				ctx.Send(message.Text(fmt.Sprintf("%v已被踢出群聊", qq)))
			}
		}
	})

	// 退出群聊
	plugin.OnRegex("^退出群聊(/d+)", leafbot.Option{
		Weight: 1,
		Block:  true,
		Rules:  []leafbot.Rule{UserSuperUser},
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
		Rules:  []leafbot.Rule{UserAdminPermission},
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
		ctx.Send(append(message.Message{}, message.Text("这个头衔很不错吧"), message.At(int64(ctx.Event.UserId))))
	})

	// 修改群名片
	plugin.OnRegex(`^修改名片.*?(\d+).*?\s(.*)`, leafbot.Option{
		Weight: 10,
		Block:  true,
		Rules:  []leafbot.Rule{UserAdminPermission},
	}).Handle(func(ctx *leafbot.Context) {
		id, err := strconv.ParseInt(ctx.State.RegexResult[1], 10, 64)
		if err != nil {
			return
		}
		ctx.SetGroupCard(ctx.Event.GroupId, int(id), ctx.State.RegexResult[2])
	})

	plugin.OnRegex(`^禁言.*?(\d+).*?\s(\d+)(.*)`, leafbot.Option{
		Weight: 10,
		Block:  true,
		Rules:  []leafbot.Rule{UserAdminPermission},
	}).Handle(func(ctx *leafbot.Context) {
		id, err := strconv.ParseInt(ctx.State.RegexResult[1], 10, 64)
		if err != nil {
			return
		}
		duration, err := strconv.ParseInt(ctx.State.RegexResult[2], 10, 64)
		if err != nil {
			return
		}
		switch ctx.State.RegexResult[3] {
		case "天", "day":
			duration = duration * 60 * 60 * 24
		case "分钟", "min", "m":
			duration = duration * 60
		case "h", "小时":
			duration = duration * 60 * 60
		default:
			duration = duration * 60
		}
		if duration >= 43200 {
			duration = 43199 // qq禁言最大时长为一个月
		}
		ctx.SetGroupBan(ctx.Event.GroupId, int(id), int(duration))
		ctx.Send(message.Text("小黑屋收留成功"))
	})

	plugin.OnRegex(`^解除禁言.*?(\d+)`, leafbot.Option{
		Weight: 10,
		Block:  true,
		Rules:  []leafbot.Rule{UserAdminPermission},
	}).Handle(func(ctx *leafbot.Context) {
		id, err := strconv.ParseInt(ctx.State.RegexResult[1], 10, 64)
		if err != nil {
			return
		}
		ctx.SetGroupBan(ctx.Event.GroupId, int(id), 0)
	})

	plugin.OnFullMatch("123", leafbot.Option{
		Weight: 1,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		log.Infoln("匹配的第一个")
	})

	plugin.OnFullMatch("123", leafbot.Option{
		Weight: 2,
		Block:  true,
	}).Handle(func(ctx *leafbot.Context) {
		log.Infoln("匹配的第二个")
	})

}

func UserAdminPermission(ctx *leafbot.Context) bool {
	if ctx.Event.Sender.Role == "owner" || ctx.Event.Sender.Role == "admin" {
		return true
	}

	if ctx.Event.UserId == leafbot.GetLeafConfig().Admin {
		return true
	}

	for _, user := range leafbot.GetLeafConfig().SuperUser {
		if user == ctx.Event.UserId {
			return true
		}
	}

	return false

}

func UserSuperUser(ctx *leafbot.Context) bool {
	for _, user := range leafbot.GetLeafConfig().SuperUser {
		if user == ctx.Event.UserId {
			return true
		}
	}

	return false
}
