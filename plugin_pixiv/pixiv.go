package plugin_pixiv

import (
	"encoding/base64"
	"strconv"

	"github.com/guonaihong/gout"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"

	pixiv2 "github.com/huoxue1/leafBotPlugin/global/pixiv"
)

func init() {
	pixiv()
}

func pixiv() {
	plugin := leafBot.NewPlugin("pixiv")
	plugin.OnCommand("添加图片", leafBot.Option{
		PluginName: "图片添加",
		Weight:     1,
		Block:      false,
		Allies:     []string{"add"},
	}).AddRule(func(event leafBot.Event, api leafBot.Api, state *leafBot.State) bool {
		for _, segment := range event.Message {
			if segment.Type == "reply" {
				state.Data["reply_id"] = segment.Data["id"]
			}
		}
		return true
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		if id, ok := state.Data["reply_id"]; ok {
			bot.GetMsg(id.(int32)).Get("message").ForEach(func(key, value gjson.Result) bool {
				if value.Get("type").String() == "image" {
					var resp []byte
					err := gout.GET(value.Get("data.url").String()).BindBody(&resp).Do()
					if err != nil {
						return false
					}
					image := Image{
						ID:      uuid.NewV4().String(),
						ImageID: 0,
						Content: base64.StdEncoding.EncodeToString(resp),
					}
					err = db.Insert("image", &image)
					if err != nil {
						return false
					}
				}
				return true
			})
		} else {
			if len(state.Args) < 1 {
				return
			}
			data, err := pixiv2.GetDataByID(state.Args[0])
			if err != nil {
				return
			}
			id, err := strconv.Atoi(state.Args[0])
			if err != nil {
				return
			}
			i := Image{
				ID:      uuid.NewV4().String(),
				ImageID: int64(id),
				Content: data,
			}
			err = db.Insert("image", &i)
			if err != nil {
				return
			}
			event.Send(message.Message{message.Text("已成功添加图片，id:" + state.Args[0]), message.Image("base64://" + data)})
		}
	})

	plugin.OnCommand("来点色图", leafBot.Option{
		PluginName: "setu",
		Weight:     1,
		Block:      false,
		Allies:     nil,
		Rules:      []leafBot.Rule{leafBot.OnlySuperUser},
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		if len(state.Args) < 1 {
			return
		}
		bot.SendGroupForwardMsg(event.GroupId, pixiv2.Search(event, state.Args[0]))
	})

	plugin.OnCommand("榜单参数", leafBot.Option{
		PluginName: "榜单",
		Weight:     1,
		Block:      false,
		Allies:     nil,
		Rules:      []leafBot.Rule{leafBot.OnlyToMe},
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		event.Send(message.Text("daily\n\nweekly\n\nmonthly\n\nrookie\n\noriginal\n\nmale\n\nfemale\n\ndaily_r18\n\nweekly_r18\n\nmale_r18\n\nfemale_r18\n\nr18g"))
	})

	plugin.OnCommand("榜单", leafBot.Option{
		PluginName: "rank",
		Weight:     1,
		Block:      false,
		Allies:     nil,
		Rules:      []leafBot.Rule{leafBot.OnlySuperUser},
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		if len(state.Args) < 1 {
			return
		}
		messages := pixiv2.GetWeek(event, state.Args[0])
		bot.SendGroupForwardMsg(event.GroupId, messages)
	})
}
