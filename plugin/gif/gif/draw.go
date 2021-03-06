// Package gif
// @Description:
package gif

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/gif"

	log "github.com/sirupsen/logrus"
	"github.com/tdf1939/img"
)

func Mo(yuan string) string {
	im := img.ImDc(yuan, 0, 0).Circle(0).Im
	mo := []*image.NRGBA{
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/mo/0.png", 0, 0).DstOver(im, 80, 80, 32, 32).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/mo/1.png", 0, 0).DstOver(im, 70, 90, 42, 22).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/mo/2.png", 0, 0).DstOver(im, 75, 85, 37, 27).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/mo/3.png", 0, 0).DstOver(im, 85, 75, 27, 37).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/mo/4.png", 0, 0).DstOver(im, 90, 70, 22, 42).Im,
	}

	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, mo))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func Cuo(yuan string) string {
	tou := img.ImDc(yuan, 110, 110).Circle(0).Im
	m1 := img.Rotate(tou, 72, 0, 0)
	m2 := img.Rotate(tou, 144, 0, 0)
	m3 := img.Rotate(tou, 216, 0, 0)
	m4 := img.Rotate(tou, 288, 0, 0)
	cuo := []*image.NRGBA{
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/cuo/0.png", 0, 0).DstOverC(tou, 0, 0, 75, 130).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/cuo/1.png", 0, 0).DstOverC(m1.Im, 0, 0, 75, 130).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/cuo/2.png", 0, 0).DstOverC(m2.Im, 0, 0, 75, 130).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/cuo/3.png", 0, 0).DstOverC(m3.Im, 0, 0, 75, 130).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/cuo/4.png", 0, 0).DstOverC(m4.Im, 0, 0, 75, 130).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(5, cuo))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func Pai(yuan string) string {
	tou := img.ImDc(yuan, 30, 30).Circle(0).Im
	pai := []*image.NRGBA{
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/pai/0.png", 0, 0).Over(tou, 0, 0, 1, 47).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/pai/1.png", 0, 0).Over(tou, 0, 0, 1, 67).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, pai))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func Chong(yuan string) string {
	tou := img.ImDc(yuan, 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/xqe/0.png", 0, 0).Over(tou, 30, 30, 15, 53).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/xqe/1.png", 0, 0).Over(tou, 30, 30, 40, 53).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, chong))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

// Qiao 敲
func Qiao(yuan string) string {
	tou := img.ImDc(yuan, 40, 40).Circle(0).Im
	qiao := []*image.NRGBA{
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/qiao/0.png", 0, 0).Over(tou, 40, 33, 57, 52).Im,
		img.ImDc("https://codechina.csdn.net/m15082717021/image/-/raw/main/qiao/1.png", 0, 0).Over(tou, 38, 36, 58, 50).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, qiao))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

// Chi 吃
func Chi(yuan string) string {
	tou := img.ImDc(yuan, 32, 32).Im
	chi := []*image.NRGBA{
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/chi/0.png`, 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/chi/1.png`, 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/chi/2.png`, 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, chi))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

// Ken 啃
func Ken(yuan string) string {
	tou := img.ImDc(yuan, 100, 100).Im
	ken := []*image.NRGBA{
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/0.png`, 0, 0).DstOver(tou, 90, 90, 105, 150).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/1.png`, 0, 0).DstOver(tou, 90, 83, 96, 172).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/2.png`, 0, 0).DstOver(tou, 90, 90, 106, 148).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/3.png`, 0, 0).DstOver(tou, 88, 88, 97, 167).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/4.png`, 0, 0).DstOver(tou, 90, 85, 89, 179).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/5.png`, 0, 0).DstOver(tou, 90, 90, 106, 151).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/6.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/7.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/8.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/9.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/10.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/11.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/12.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/13.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/14.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/ken/15.png`, 0, 0).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, ken))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}

//丢
func Diu(yuan string) string {
	tou := img.ImDc(yuan, 0, 0).Circle(0).Im
	diu := []*image.NRGBA{
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main//diu/0.png`, 0, 0).Over(tou, 32, 32, 108, 36).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/1.png`, 0, 0).Over(tou, 32, 32, 122, 36).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/2.png`, 0, 0).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/3.png`, 0, 0).Over(tou, 123, 123, 19, 129).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/4.png`, 0, 0).Over(tou, 185, 185, -50, 200).Over(tou, 33, 33, 289, 70).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/5.png`, 0, 0).Over(tou, 32, 32, 280, 73).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/6.png`, 0, 0).Over(tou, 35, 35, 259, 31).Im,
		img.ImDc(`https://codechina.csdn.net/m15082717021/image/-/raw/main/diu/7.png`, 0, 0).Over(tou, 175, 175, -50, 220).Im,
	}
	var result []byte
	buffer := bytes.NewBuffer(result)
	err := gif.EncodeAll(buffer, img.AndGif(1, diu))
	if err != nil {
		log.Panicln("生成图片出现错误")
	}
	return "base64://" + base64.StdEncoding.EncodeToString(buffer.Bytes())
}
