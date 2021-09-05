package main

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/utils"
	"github.com/huoxue1/leafBotPlugin/plugin_gif"
	log "github.com/sirupsen/logrus"

	// 导入插件
	_ "github.com/huoxue1/leafBotPlugin/pluginBlackList"
	_ "github.com/huoxue1/leafBotPlugin/pluginDayImage"
	_ "github.com/huoxue1/leafBotPlugin/pluginFlashImage"
	_ "github.com/huoxue1/leafBotPlugin/pluginGithub"
	_ "github.com/huoxue1/leafBotPlugin/pluginGroupManager"
	_ "github.com/huoxue1/leafBotPlugin/pluginHelp"
	_ "github.com/huoxue1/leafBotPlugin/pluginMusic"
	_ "github.com/huoxue1/leafBotPlugin/pluginOcr"
	_ "github.com/huoxue1/leafBotPlugin/pluginParseMessage"
	_ "github.com/huoxue1/leafBotPlugin/pluginPoke"
	_ "github.com/huoxue1/leafBotPlugin/pluginQrCode"
	_ "github.com/huoxue1/leafBotPlugin/pluginSearchImage"
	_ "github.com/huoxue1/leafBotPlugin/pluginSmallTalk"
	_ "github.com/huoxue1/leafBotPlugin/pluginTranslate"
	_ "github.com/huoxue1/leafBotPlugin/pluginWebsite"
	_ "github.com/huoxue1/leafBotPlugin/pluginWeibo"
)

var VERSION = "UnKnow"

func init() {
	leafBot.InitPluginManager()
	plugin_gif.MoInit()
	go utils.PwInit()
}

func main() {
	go leafBot.InitWindow()
	log.Infoln("当前版本>>>>>>>>  " + VERSION)
	leafBot.InitBots()
}
