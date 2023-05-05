package flash

import (
	"fmt"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
)

func init() {
	plugin := leafbot.NewPlugin("闪照")
	plugin.OnMessage("private", leafbot.Option{
		Weight: 0,
		Block:  true,
		Allies: nil,
		Rules: []leafbot.Rule{func(ctx *leafbot.Context) bool {
			for _, segment := range ctx.Event.Message {
				if segment.Type == "image" {
					return true
				}
			}
			return false
		}},
	}).Handle(func(ctx *leafbot.Context) {
		img := ctx.Event.Message[0]
		delete(img.Data, "type")
		ctx.SendGroupMsg(972264701, []message.MessageSegment{message.Text(fmt.Sprintf("来自用户%d的闪照", ctx.UserID)), img})
	})
}
