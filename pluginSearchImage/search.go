package pluginSearchImage

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"net/http"
)

func init() {
	InitImage()
}

func InitImage() {

	plugin := leafBot.NewPlugin("图片搜索")
	plugin.SetHelp(map[string]string{
		"搜图": "图片搜索",
	})
	plugin.OnCommand("搜图").SetPluginName("搜图").SetWeight(10).SetBlock(false).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		images, err := SearchImage(state.Args[0])
		if err != nil {
			event.Send(message.Text("接口报错了" + err.Error()))
			return
		}
		mess := message.Message{}
		for _, image := range images {
			mess = append(mess, message.Image(image))
			//mess = append(mess, message.CustomNode(event.Sender.NickName, int64(event.UserId), "[CQ:image,file="+image+"]"))

		}

		event.Send(mess)
		//bot.SendGroupForwardMsg(event.GroupId, mess)
	})
}

func SearchImage(ketWord string) ([]string, error) {
	resp, err := http.Get("https://pixiv.kurocore.com/illust?keyword=" + ketWord)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var list []string
	doc.Find("div.illust-image a.cover img").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("data-original")
		list = append(list, "http:"+link)
	})
	return list, err
}
