// Package plugin_gif
// 该插件来源于 https://github.com/tdf1939/ZeroBot-Plugin-Gif
///*
package plugin_gif

import (
	"fmt"
	"regexp"

	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"

	"github.com/huoxue1/leafBotPlugin/plugin_gif/gif"
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

func MoInit() {
	plugin := leafBot.NewPlugin("搞笑gif")
	plugin.SetHelp(map[string]string{
		"摸": "",
		"搓": "",
		"冲": "",
		"拍": "",
		"敲": "",
		"吃": "",
		"啃": "",
		"丢": ""})
	plugin.OnMessage("group").SetWeight(10).SetPluginName("gif").AddRule(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) bool {
		for s := range m {
			if event.Message[0].Type == "text" && event.Message[0].Data["text"] == s {
				state.Data["type"] = event.Message[0].Data["text"]
				for _, segment := range event.Message {
					if segment.Type == "at" {
						state.Data["data"] = segment.Data["qq"]
					} else if segment.Type == "image" {
						state.Data["image"] = segment.Data["url"]
					}
				}
			}
		}
		return true
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		f := m[state.Data["type"].(string)]
		link := ""
		data, ok := state.Data["data"]
		data1, ok1 := state.Data["image"]
		switch {
		case ok:
			link = fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%v&s=100", data)
		case ok1:
			link = data1.(string)
		default:
			compile := regexp.MustCompile(`\d+`)
			if compile.MatchString(event.GetPlainText()) {
				link = fmt.Sprintf("http://q1.qlogo.cn/g?b=qq&nk=%v&s=100", compile.FindStringSubmatch(event.GetPlainText())[1])
			}
		}
		event.Send(message.Image(f(link)))
	})
}
