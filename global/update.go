package global

import (
	"fmt"
	"github.com/guonaihong/gout"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
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
	var content []byte
	switch runtime.GOOS {
	case "windows":
		err := gout.GET(fmt.Sprintf("https://github.com/huoxue1/leafBotPlugin/releases/download/%v/leafBotPlugin_windows_amd64.exe", version)).BindBody(&content).Do()
		if err != nil {
			return err
		}
	case "linux":
		err := gout.GET(fmt.Sprintf("https://github.com/huoxue1/leafBotPlugin/releases/download/%v/leafBotPlugin_windows_amd64.exe", version)).BindBody(&content).Do()
		if err != nil {
			return err
		}
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
