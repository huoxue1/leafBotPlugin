package global

import (
	"fmt"
	"github.com/guonaihong/gout"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Update() error {
	var data string
	err := gout.GET("https://api.github.com/repos/huoxue1/leafBotPlugin/releases/latest").BindBody(&data).Do()
	if err != nil {
		return err
	}
	version := gjson.Get(data, "tag_name").Str
	log.Infoln("正在开始更新——————————————")
	switch runtime.GOOS {
	case "windows":
		response, err := gout.GET(fmt.Sprintf("https://github.com/huoxue1/leafBotPlugin/releases/download/%v/leafBotPlugin_windows_amd64.exe", version)).Response()
		if err != nil {
			return err
		}
		file, err := os.OpenFile("./leafBotPlugin_windows_"+version+"_amd64.exe", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer func() {
			file.Close()
			response.Body.Close()
		}()
		bar := &Bar{}
		bar.NewOption(0, response.ContentLength, response.Body)
		_, err = io.Copy(file, bar)
		if err != nil {
			return err
		}
		return err
	case "linux":
		response, err := gout.GET(fmt.Sprintf("https://github.com/huoxue1/leafBotPlugin/releases/download/%v/leafBotPlugin_linux_amd64.exe", version)).Response()
		if err != nil {
			return err
		}
		file, err := os.OpenFile("./leafBotPlugin_linux_"+version+"_amd64.exe", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer func() {
			file.Close()
			response.Body.Close()
		}()
		bar := &Bar{}
		bar.NewOption(0, response.ContentLength, response.Body)
		_, err = io.Copy(file, bar)
		if err != nil {
			return err
		}
		return err
	}
	return err
}

func GetLastVersion() (string, error) {
	var data string
	err := gout.GET("https://api.github.com/repos/huoxue1/leafBotPlugin/releases/latest").BindBody(&data).Do()
	if err != nil {
		return "", err
	}
	return gjson.Get(data, "tag_name").Str, err
}

func CheckVersion(oldVersion, newVersion string) bool {
	if oldVersion == "UnKnow" {
		log.Infoln("使用action版本或者自编译版本")
		return false
	}
	oldVersion = strings.ReplaceAll(oldVersion, ".", "")
	oldVersion = strings.ReplaceAll(oldVersion, "v", "")
	newVersion = strings.ReplaceAll(newVersion, ".", "")
	newVersion = strings.ReplaceAll(newVersion, "v", "")
	old, err := strconv.Atoi(oldVersion)
	newV, err := strconv.Atoi(newVersion)
	if err != nil {
		return false
	}

	return newV > old
}
