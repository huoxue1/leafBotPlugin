package plugin_course

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/huoxue1/gg"
	"github.com/huoxue1/leafBot"
)

func draw(courses []Course) string {
	file, err := os.OpenFile("./config/model.png", os.O_RDONLY, 0666)
	if err != nil {
		return ""
	}
	im, _, err := image.Decode(file)
	if err != nil {
		return ""
	}
	d := gg.NewContextForImage(im)
	err = d.LoadFontFromBytes(leafBot.GetFont(), 15)
	if err != nil {
		return ""
	}
	//d.DrawRectangle(185, 58, 205, 89)
	//d.SetRGB255(100, 100, 0)
	//d.Fill()
	f := 0.0
	id := 0
	for _, cours := range courses {
		if cours.ID != id {
			d.SetRGB255(getRandomColor())
		}
		id = cours.ID
		f = float64(58 + ((cours.Time - 1) * 89))
		f += 3
		d.DrawRectangle(185, f, 308, 85)
		d.Fill()
	}
	f = 0
	d.SetRGB255(255, 255, 255)
	for _, cours := range courses {
		f = float64(58 + ((cours.Time - 1) * 89))
		f += 3
		d.DrawString(cours.Name, 185, f+20)
		d.DrawString(cours.Teacher, 185, f+50)
		d.DrawString(cours.Location, 185, f+80)
	}
	var result []byte
	reader := bytes.NewBuffer(result)
	err = d.EncodePNG(reader)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(reader.Bytes())
}

func getRandomColor() (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(200), rand.Intn(235), rand.Intn(100)
}
