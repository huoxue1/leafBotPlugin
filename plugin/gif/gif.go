package gif

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"

	"github.com/huoxue1/fan/plugin/gif/gif"
)

var (
	m = map[string]func(string) string{
		"摸": gif.Mo,
		"搓": gif.Cuo,
		"冲": gif.Chong,
		"拍": gif.Pai,
		"敲": gif.Qiao,
		"吃": gif.Chi,
		"啃": gif.Ken,
		"丢": gif.Diu,
	}
)

func init() {
	plugin := leafbot.NewPlugin("gif")
	plugin.OnMessage("", leafbot.Option{
		Weight: 10,
		Block:  false,
		Rules:  nil,
	}).Handle(func(ctx *leafbot.Context) {
		f, ok := m[ctx.Event.Message[0].Data["text"]]
		if !ok {
			return
		}
		users := ctx.GetAtUsers()
		if len(users) > 0 {
			ctx.Send(message.Message{message.Reply(int64(ctx.Event.MessageID)), message.Image("base64://" + f(fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%v&s=100", users[0])))})
		} else {
			if len(ctx.GetImages()) > 0 {
				ctx.Send(message.Message{message.Reply(int64(ctx.Event.MessageID)), message.Image(f(ctx.GetImages()[0].Data["url"]))})
			} else {
				messages := strings.Split(ctx.Event.Message.ExtractPlainText(), " ")
				if len(messages) > 1 {
					id, err := strconv.ParseInt(messages[1], 10, 64)
					if err != nil {
						return
					}
					ctx.Send(message.Message{message.Reply(int64(ctx.Event.MessageID)), message.Image("base64://" + f(fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%v&s=100", id)))})
				}
			}
		}
	})
}
