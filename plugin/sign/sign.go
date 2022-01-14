package sign

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
	log "github.com/sirupsen/logrus"
)

func init() {
	do()
}

func do() {
	plugin := leafbot.NewPlugin("签到")
	plugin.OnFullMatchGroup("签到", leafbot.Option{
		Weight: 10,
		Block:  true,
		Allies: []string{"打卡"},
	}).Handle(func(ctx *leafbot.Context) {
		sign(ctx)
	})

	plugin.OnFullMatchGroup("积分查询", leafbot.Option{
		Weight: 10,
		Block:  true,
		Allies: []string{"我的积分"},
	}).Handle(func(ctx *leafbot.Context) {
		s := new(Sign)
		s.QQ = int64(ctx.Event.UserId)
		err := query(s)
		if err != nil {
			return
		}
		ctx.Send([]message.MessageSegment{
			message.Text(fmt.Sprintf("你当前积分为%d,已连续签到%d天", s.Fraction, s.ContinueSign)),
			message.At(int64(ctx.Event.UserId))})
	})
}

func sign(ctx *leafbot.Context) {
	s := new(Sign)
	s.QQ = int64(ctx.Event.UserId)
	haveContent(int64(ctx.Event.UserId))
	err := query(s)
	if err != nil {
		log.Errorln("查询个人信息失败")
		log.Errorln(err.Error())
		return
	}
	if isToday(s.LastSign) {
		ctx.Send(message.Text("今日已经签过到了额，请明日再来吧"))
		return
	}
	n := rand.Int63n(3) + 1
	s.Fraction += int64(n)
	if isYesterday(s.LastSign) {
		s.ContinueSign++
		err := update(*s)
		if err != nil {
			log.Errorln("更新用户信息出现了错误")
			return
		}
	}
	s.LastSign = time.Now().Unix()
	update(*s)
	ctx.Send(append(message.Message{},
		//message.Image(fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%d&s=640", ctx.Event.UserId)).Add("cache", 0),
		message.Text(fmt.Sprintf("恭喜你，签到成功,积分增加%d,当前共有积分%d,已连续签到%d天", n, s.Fraction, s.ContinueSign)),
		message.At(int64(ctx.Event.UserId))))

}

func isToday(times int64) bool {
	signTime := time.Unix(times, 0)
	now := time.Now()
	if now.Year() == signTime.Year() && now.Month() == signTime.Month() && now.Day() == signTime.Day() {
		return true
	}
	return false
}

func isYesterday(times int64) bool {
	signTime := time.Unix(times, 0)
	duration, _ := time.ParseDuration("-1d")
	yesterday := time.Now().Add(duration)
	if yesterday.Year() == signTime.Year() && yesterday.Month() == signTime.Month() && yesterday.Day() == signTime.Day() {
		return true
	}
	return false
}
