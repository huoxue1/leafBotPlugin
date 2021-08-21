package pluginWeibo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"github.com/huoxue1/leafBot/utils"
	"github.com/mxschmitt/playwright-go"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func init() {
	leafBot.OnCommand("/weibo").
		AddAllies("热搜").
		SetWeight(10).
		SetBlock(false).
		SetPluginName("微博热搜").
		AddHandle(weiBoHandle)
}

func weiBoHandle(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
	if leafBot.DefaultConfig.Plugins.AlApiToken == "" {
		bot.Send(event, message.Text("未检测到alapitoken，请联系bot管理员为其配置。\n申请地址https://admin.alapi.cn/api_manager/token_manager"))
		return
	}
	if len(state.Args) < 1 {
		draw(50)
		srcByte, err := ioutil.ReadFile("./plugins/weibo/weibo.png")
		if err != nil {
			log.Fatal(err)
		}

		res := base64.StdEncoding.EncodeToString(srcByte)

		bot.Send(event, message.Image("base64://"+res))
	} else {
		limit, err := strconv.Atoi(state.Args[0])
		if err != nil {
			return
		}
		if limit > 50 || limit < 0 {
			bot.Send(event, []message.MessageSegment{message.Text("非法参数"), message.At(int64(event.UserId))})
			return
		}
		//draw(limit)
		api, err := getDataAlApi(50)
		if err != nil {
			bot.Send(event, message.Text("api获取错误"+err.Error()))
			return
		}
		bot.Send(event, message.Text("downloading image ......"))
		data, err := utils.GetPWScreen(fmt.Sprintf("https://s.weibo.com/weibo?q=%v&Refer=top", api.Data[limit-1].HotWord), "android")
		if err != nil {
			bot.Send(event, err.Error())
			return
		}
		bot.Send(event, message.Image("base64://"+base64.StdEncoding.EncodeToString(data)))
	}

	//getWeibo(0)
}

func getWeibo(url string) []byte {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()

	defer func() {
		page.Close()
		pw.Stop()
	}()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	data, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String("pla.png"),
		FullPage: playwright.Bool(true),
	})
	return data

}

func draw(limit int) {
	context := gg.NewContext(900, 100*(limit+1))
	context.SetRGB255(255, 255, 0)
	context.DrawRectangle(0, 0, 900, float64(100*(limit+1)))
	//weibo, err := getData()
	weibo, err := getDataAlApi(limit)
	context.Fill()
	if err := context.LoadFontFace("./config/NotoSansBold.ttf", 40); err != nil {
		log.Debugln(err)
	}
	context.SetRGB255(0, 0, 0)
	fmt.Println(weibo)
	if err != nil {
		return
	}
	//for i := 0; i < limit; i++ {
	//	fmt.Println(weibo.Data[i].Name)
	//	context.DrawString(strconv.Itoa(i+1)+"："+weibo.Data[i].Name, 0, float64(100*(i+1)))
	//}
	for i, datum := range weibo.Data {
		context.DrawString(strconv.Itoa(i+1)+"："+datum.HotWord+"  "+datum.HotWordNum, 0, float64(100*(i+1)))
	}
	err = context.SavePNG("./plugins/weibo/weibo.png")
	if err != nil {
		log.Debugln("图片保存失败")
	}
}

func getDataAlApi(num int) (AlApi, error) {
	resp, err := http.Get("https://v2.alapi.cn/api/new/wbtop?token=" + leafBot.DefaultConfig.Plugins.AlApiToken + "&num=" + strconv.Itoa(num))
	if err != nil {
		return AlApi{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AlApi{}, err
	}
	weibo := AlApi{}
	err = json.Unmarshal(data, &weibo)
	return weibo, err
}

func getData() (Weibo, error) {
	resp, err := http.Get("https://api.hmister.cn/weibo/")
	if err != nil {
		return Weibo{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weibo{}, err
	}
	weibo := Weibo{}
	err = json.Unmarshal(data, &weibo)
	return weibo, err
}
