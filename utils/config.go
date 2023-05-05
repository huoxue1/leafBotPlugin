package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Fiction struct {
		RemoteAddress string `json:"remote_address" yaml:"remote_address"`
		UploadSecret  string `json:"upload_secret" yaml:"upload_secret"`
	} `json:"fiction" yaml:"fiction"`

	Model struct {
		Driver string `json:"driver" yaml:"driver"`
		Dsl    string `json:"dsl" yaml:"dsl"`
	} `json:"model" yaml:"model"`

	Bot struct {
		NickName  string  `json:"nick_name" yaml:"nick_name"`
		SuperUser []int64 `json:"super_user" yaml:"super_user"`
	} `json:"bot" yaml:"bot"`
}

var (
	config Config
)

func init() {
	initConfig()
}

func initConfig() {
	data, err := os.ReadFile("leafbot.yml")
	if err != nil {
		log.Panicln("配置文件leafbot.yml读取失败" + err.Error())
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Errorln("配置信息解析失败" + err.Error())
		return
	}
	log.Infoln("已加载配置文件")
	log.Infoln(config)
}

func GetConfig() *Config {
	return &config
}
