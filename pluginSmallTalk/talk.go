package pluginSmallTalk

import (
	"github.com/guonaihong/gout"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"net/url"
	"strings"
)

func init() {
	InitSmallTalk()
}

func InitSmallTalk() {
	leafBot.OnCommand("开启闲聊").SetPluginName("闲聊开启插件").SetBlock(false).SetWeight(1).AddRule(leafBot.OnlyToMe).AddHandle(
		func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
			bot.Send(event, message.Text("闲聊开启成功,输入关闭闲聊即可退出闲聊模式"))
			var data string
			for data != "关闭闲聊" {
				oneEvent, err := bot.GetOneEvent(func(event1 leafBot.Event, bot *leafBot.Bot, state *leafBot.State) bool {
					if event1.GroupId == event.GroupId && event1.UserId == event.UserId {
						return true
					}
					return false
				})
				if err != nil {
					bot.Send(event, message.Text("获取数据超时，提前退出闲聊模式"))
					return
				}
				data = oneEvent.Message.ExtractPlainText()
				talk, err := getTalk(data)
				if err != nil {
					bot.Send(event, "闲聊api出现故障\n"+err.Error())
					return
				}
				bot.Send(event, talk.Content)
			}
		})
}

type res struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

func getTalk(data string) (res, error) {
	encodeData := url.QueryEscape(data)
	r := new(res)
	err := gout.GET("http://api.qingyunke.com/api.php?key=free&appid=0&msg=" + encodeData).BindJSON(r).Err
	if err != nil {
		return *r, err
	}
	r.Content = strings.ReplaceAll(r.Content, "{br}", "\n")
	return *r, err
}
