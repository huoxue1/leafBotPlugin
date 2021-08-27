package gif

import (
	"encoding/base64"
	"github.com/tdf1939/img"
	"image"
	"io/ioutil"
)

func Mo(yuan string) string {

	im := img.ImDc(yuan, 0, 0).Circle(0).Im
	mo := []*image.NRGBA{
		img.ImDc("https://specialblog.link/img/202108272129261.png", 0, 0).DstOver(im, 80, 80, 32, 32).Im,
		img.ImDc("https://specialblog.link/img/202108272129471.png", 0, 0).DstOver(im, 70, 90, 42, 22).Im,
		img.ImDc("https://specialblog.link/img/202108272130562.png", 0, 0).DstOver(im, 75, 85, 37, 27).Im,
		img.ImDc("https://specialblog.link/img/202108272130451.png", 0, 0).DstOver(im, 85, 75, 27, 37).Im,
		img.ImDc("https://specialblog.link/img/202108272130890.png", 0, 0).DstOver(im, 90, 70, 22, 42).Im,
	}
	img.SaveGif(img.AndGif(1, mo), "./tmp/mo.gif")
	file, _ := ioutil.ReadFile("./tmp/mo.gif")
	return "base64://" + base64.StdEncoding.EncodeToString(file)
}

func Cuo(yuan string) string {
	tou := img.ImDc(yuan, 110, 110).Circle(0).Im
	m1 := img.Rotate(tou, 72, 0, 0)
	m2 := img.Rotate(tou, 144, 0, 0)
	m3 := img.Rotate(tou, 216, 0, 0)
	m4 := img.Rotate(tou, 288, 0, 0)
	cuo := []*image.NRGBA{
		img.ImDc("https://specialblog.link/img/202108272144822.png", 0, 0).DstOverC(tou, 0, 0, 75, 130).Im,
		img.ImDc("https://specialblog.link/img/202108272144885.png", 0, 0).DstOverC(m1.Im, 0, 0, 75, 130).Im,
		img.ImDc("https://specialblog.link/img/202108272144344.png", 0, 0).DstOverC(m2.Im, 0, 0, 75, 130).Im,
		img.ImDc("https://specialblog.link/img/202108272145718.png", 0, 0).DstOverC(m3.Im, 0, 0, 75, 130).Im,
		img.ImDc("https://specialblog.link/img/202108272145323.png", 0, 0).DstOverC(m4.Im, 0, 0, 75, 130).Im,
	}
	img.SaveGif(img.AndGif(5, cuo), "./tmp/cuo.gif")
	file, _ := ioutil.ReadFile("./tmp/cuo.gif")
	return "base64://" + base64.StdEncoding.EncodeToString(file)
}

func Pai(yuan string) string {
	tou := img.ImDc(yuan, 30, 30).Circle(0).Im
	pai := []*image.NRGBA{
		img.ImDc("https://specialblog.link/img/202108272203655.png", 0, 0).Over(tou, 0, 0, 1, 47).Im,
		img.ImDc("https://specialblog.link/img/202108272203063.png", 0, 0).Over(tou, 0, 0, 1, 67).Im,
	}
	img.SaveGif(img.AndGif(1, pai), "./tmp/pai.gif")
	file, _ := ioutil.ReadFile("./tmp/pai.gif")
	return "base64://" + base64.StdEncoding.EncodeToString(file)
}

func Chong(yuan string) string {
	tou := img.ImDc(yuan, 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.ImDc("https://specialblog.link/img/202108272205152.png", 0, 0).Over(tou, 30, 30, 15, 53).Im,
		img.ImDc("https://specialblog.link/img/202108272205307.png", 0, 0).Over(tou, 30, 30, 40, 53).Im,
	}
	img.SaveGif(img.AndGif(1, chong), "./tmp/chong.gif")
	file, _ := ioutil.ReadFile("./tmp/chong.gif")
	return "base64://" + base64.StdEncoding.EncodeToString(file)

}
