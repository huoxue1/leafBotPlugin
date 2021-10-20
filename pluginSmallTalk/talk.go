package pluginSmallTalk

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	log "github.com/sirupsen/logrus"
)

func init() {
	go InitSmallTalk()
}

func InitSmallTalk() {
	plugin := leafBot.NewPlugin("智障闲聊")
	plugin.SetHelp(map[string]string{"开启闲聊": "进入闲聊模式"})
	plugin.OnCommand("开启闲聊").SetPluginName("闲聊开启插件").SetBlock(false).SetWeight(1).AddRule(leafBot.OnlyToMe).AddHandle(
		func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
			event.Send(message.Text("闲聊开启成功,输入关闭闲聊即可退出闲聊模式"))
			var data string
			for data != "关闭闲聊" {
				oneEvent, err := event.GetOneEvent(func(event1 leafBot.Event, bot leafBot.Api, state *leafBot.State) bool {
					if event1.GroupId == event.GroupId && event1.UserId == event.UserId {
						return true
					}
					return false
				})
				if err != nil {
					event.Send(message.Text("获取数据超时，提前退出闲聊模式"))
					return
				}
				data = oneEvent.Message.ExtractPlainText()
				talk, err := getTalk(data)
				if err != nil {
					event.Send("闲聊api出现故障\n" + err.Error())
					return
				}
				event.Send(talk.Content)
			}
		})
}

type res struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

func getTalk(data string) (res, error) {
	encodeData := url.QueryEscape(data)
	r := res{}
	resp, err := http.Get("http://api.qingyunke.com/api.php?key=free&appid=0&msg=" + encodeData)
	if err != nil {
		return res{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return res{}, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return res{}, err
	}
	r.Content = strings.ReplaceAll(r.Content, "{br}", "\n")
	log.Infoln(r.Content)
	return r, err
}
