package plugin_pixiv

import "github.com/huoxue1/leafBotPlugin/global/data"

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
	err := db.Create("image", &Image{})
	if err != nil {
		return
	}
}
