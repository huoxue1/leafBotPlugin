package pixiv

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/NateScarlet/pixiv/pkg/artwork"
	client2 "github.com/NateScarlet/pixiv/pkg/client"
	"github.com/guonaihong/gout"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
)

var (
	c         *http.Client
	client    *client2.Client
	PHPSESSID = ""
	headers   = map[string]string{
		"User-Agent": client2.DefaultUserAgent,
		"Referer":    "https://www.pixiv.net/",
	}
)

func init() {
	id, ok := leafBot.DefaultConfig.Datas["pixiv_id"].(string)
	if ok {
		PHPSESSID = id
	}
	log.Infoln("ph:::     " + PHPSESSID)
	parse, _ := url.Parse("http://127.0.0.1:7890")
	c = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(parse),
		},
	}

	client = &client2.Client{
		ServerURL: "",
		Client:    *c,
	}
	client.SetDefaultHeader("User-Agent", client2.DefaultUserAgent)
	client.SetPHPSESSID(PHPSESSID)
}

// Search
/**
 * @Description:
 * @param keyWords
 * example
 */
func Search(event leafBot.Event, keyWords string) message.Message {
	ctx := context.Background()
	ctx = client2.With(ctx, client)
	result, err := artwork.Search(ctx, keyWords, artwork.SearchOptionContentRating(artwork.ContentRatingR18))
	if err != nil {
		log.Errorln(err.Error())
	}
	d := gout.New(c)
	m := message.Message{}
	result.ForEach(func(key gjson.Result, value gjson.Result) bool {
		text := "ID: " + value.Get("id").String() + "\ntitle: " + value.Get("title").String() + "\ndescription:" + value.Get("description").String() + "\n"
		var resp []byte
		err := d.GET(value.Get("url").String()).BindBody(&resp).SetHeader(headers).Do()
		if err != nil {
			log.Errorln(err.Error())
			return true
		}
		m = append(m, message.CustomNode(event.Sender.NickName, int64(event.UserId), fmt.Sprintf(text+"[CQ:image,file=base64://%v]", base64.StdEncoding.EncodeToString(resp))))
		return true
	})
	log.Debugln(m)
	return m
}

// GetDataByID
/**
 * @Description:
 * @param id
 * example
 */
func GetDataByID(id string) (string, error) {
	ctx := context.Background()
	ctx = client2.With(ctx, client)
	i := &artwork.Artwork{ID: id}
	err := i.Fetch(ctx)
	if err != nil {
		return "", err
	}
	d := gout.New(c)
	var resp []byte
	err = d.GET(i.Image.Original).BindBody(&resp).SetHeader(headers).Do()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(resp), err
}

// GetWeek
/**
* @Description:
* @param event
* @return message.Message
* example
 */
/*
daily

weekly

monthly

rookie

original

male

female

daily_r18

weekly_r18

male_r18

female_r18

r18g
*/
func GetWeek(event leafBot.Event, model string) message.Message {
	ctx := context.Background()
	ctx = client2.With(ctx, client)
	r := &artwork.Rank{Mode: model}
	err := r.Fetch(ctx)
	if err != nil {
		return nil
	}
	d := gout.New(c)
	m := message.Message{}
	for _, item := range r.Items {
		//if len(m) > 10 {
		//	break
		//}

		fmt.Println(item.JSON.Get("url"))
		var resp []byte
		err := d.GET(item.Image.Original).BindBody(&resp).SetHeader(headers).Do()
		if err != nil {
			log.Errorln(err.Error())
			return nil
		}
		text := "ID: " + item.ID + "\nauthor: " + item.Author.Name + "\ntitle: " + item.Title + "\ndescription:" + item.Description + "\n"

		m = append(m, message.CustomNode(event.Sender.NickName, int64(event.UserId), fmt.Sprintf(text+"[CQ:image,file=base64:///%v]", base64.StdEncoding.EncodeToString(resp))))
		// m = append(m, mess)
	}
	return m
}
