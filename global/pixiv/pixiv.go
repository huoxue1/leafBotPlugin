package pixiv

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/NateScarlet/pixiv/pkg/artwork"
	client2 "github.com/NateScarlet/pixiv/pkg/client"
	"github.com/guonaihong/gout"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
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
	go initPixiv()
}

func initPixiv() {
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
	datas := result.Artworks()
	d := gout.New(c)
	m := message.Message{}
	var lock sync.Mutex
	var wait sync.WaitGroup
	wait.Add(len(result.Artworks()))

	for _, data := range datas {
		go func(artwork2 artwork.Artwork) {
			defer wait.Done()
			defer lock.Unlock()
			defer log.Infoln("下载成功", artwork2.ID)
			text := "ID: " + artwork2.ID + "\ntitle: " + artwork2.Title + "\ndescription:" + artwork2.Description + "\n"
			var resp []byte
			err := d.GET(artwork2.Image.Thumb).BindBody(&resp).SetHeader(headers).Do()
			if err != nil {
				log.Errorln(err.Error())
				return
			}
			lock.Lock()
			m = append(m, message.CustomNode(event.Sender.NickName, int64(event.UserId), fmt.Sprintf(text+"[CQ:image,file=base64://%v]", base64.StdEncoding.EncodeToString(resp))))
		}(data)
	}
	wait.Wait()
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
	var lock sync.Mutex
	var wait sync.WaitGroup
	wait.Add(len(r.Items))
	for _, item := range r.Items {
		//if len(m) > 10 {
		//	break
		//}
		go func(rankItem artwork.RankItem) {
			defer wait.Done()
			defer lock.Unlock()
			defer log.Infoln("下载成功", rankItem.ID)
			var resp []byte
			err := d.GET(rankItem.Image.Regular).BindBody(&resp).SetHeader(headers).Do()
			if err != nil {
				log.Errorln(err.Error())
				return
			}
			text := "ID: " + rankItem.ID + "\nauthor: " + rankItem.Author.Name + "\ntitle: " + rankItem.Title + "\ndescription:" + rankItem.Description + "\n"
			lock.Lock()
			m = append(m, message.CustomNode(event.Sender.NickName, int64(event.UserId), fmt.Sprintf(text+"[CQ:image,file=base64:///%v]", base64.StdEncoding.EncodeToString(resp))))

			// m = append(m, mess)
		}(item)
	}
	wait.Wait()
	return m
}
