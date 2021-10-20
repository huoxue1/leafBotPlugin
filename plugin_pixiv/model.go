package plugin_pixiv

import (
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/leafBotPlugin/global/data"
)

type Image struct {
	ID      string `db:"id"`
	ImageID int64  `db:"image_id"`
	Content string `db:"content"`
}

var (
	db = data.Sqlite{
		DB:     nil,
		DBPath: "./tmp/data.db",
	}
)

func init() {
	go Create()
}

func Create() {
	err := db.Create("image", &Image{})
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}
