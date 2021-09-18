package main

import (
	"flag"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_ws_driver"
	"github.com/huoxue1/leafBotPlugin/global"
	log "github.com/sirupsen/logrus"
	"os"
	// 导入插件
	_ "github.com/huoxue1/leafBotPlugin/pluginBlackList"
	_ "github.com/huoxue1/leafBotPlugin/pluginDayImage"
	_ "github.com/huoxue1/leafBotPlugin/pluginFlashImage"
	_ "github.com/huoxue1/leafBotPlugin/pluginGithub"
	_ "github.com/huoxue1/leafBotPlugin/pluginGroupManager"
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
	_ "github.com/huoxue1/leafBotPlugin/plugin_fakeMsg"
	_ "github.com/huoxue1/leafBotPlugin/plugin_friendRequest"
	_ "github.com/huoxue1/leafBotPlugin/plugin_localStorage"

	"github.com/huoxue1/leafBotPlugin/plugin_gif"
)

var VERSION = "UnKnow"

func init() {
	plugin_gif.MoInit()

}

func main() {
	log.Infoln("当前版本------->  " + VERSION)
	version, err := global.GetLastVersion()
	if err != nil {
		log.Errorln("检查版本失败" + err.Error())
	}
	checkVersion := global.CheckVersion(VERSION, version)
	if checkVersion {
		log.Infoln("检测到新版本" + version + "，输入--update即可自动更新")
	}
	var update bool
	flag.BoolVar(&update, "update", false, "是否更新")
	flag.Parse()
	if update {
		if checkVersion {
			err := global.Update()
			if err != nil {
				log.Errorln("检查更新失败")
			}
		} else {
			log.Warning("未检测到版本更新")
		}
		os.Exit(3)
	}
	driver := cqhttp_ws_driver.NewDriver()
	leafBot.LoadDriver(driver)
	leafBot.InitBots()
	driver.Run()
}
