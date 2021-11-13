package main

import (
	"flag"
	"os"

	"github.com/huoxue1/leafBot"                          // leafBot依赖
	"github.com/huoxue1/leafBot/cqhttp_reverse_ws_driver" // 驱动插件，cqhttp反向链接
	log "github.com/sirupsen/logrus"                      // 日志依赖

	"github.com/huoxue1/leafBotPlugin/global" // 全局工具类不
	// 导入插件
	_ "github.com/huoxue1/leafBotPlugin/pluginBlackList"      // 黑名单插件
	_ "github.com/huoxue1/leafBotPlugin/pluginDayImage"       // 每日一图插件
	_ "github.com/huoxue1/leafBotPlugin/pluginFlashImage"     // 闪照拦截插件
	_ "github.com/huoxue1/leafBotPlugin/pluginGithub"         // github查询插件
	_ "github.com/huoxue1/leafBotPlugin/pluginGroupManager"   // 群管插件
	_ "github.com/huoxue1/leafBotPlugin/pluginMusic"          // 点歌插件
	_ "github.com/huoxue1/leafBotPlugin/pluginOcr"            // ocr图像识别插件
	_ "github.com/huoxue1/leafBotPlugin/pluginParseMessage"   // 特殊消息解析插件
	_ "github.com/huoxue1/leafBotPlugin/pluginPoke"           // 戳一戳获取服务器状态插件
	_ "github.com/huoxue1/leafBotPlugin/pluginQrCode"         // 验证码生成插件
	_ "github.com/huoxue1/leafBotPlugin/pluginSearchImage"    // 图片搜索插件
	_ "github.com/huoxue1/leafBotPlugin/pluginSmallTalk"      // 闲聊插件，使用青云客插件
	_ "github.com/huoxue1/leafBotPlugin/pluginTranslate"      // 翻译插件
	_ "github.com/huoxue1/leafBotPlugin/pluginWebsite"        // 网页截图插件
	_ "github.com/huoxue1/leafBotPlugin/pluginWeibo"          // 微博热搜获取插件
	_ "github.com/huoxue1/leafBotPlugin/plugin_course"        // 课程表插件
	_ "github.com/huoxue1/leafBotPlugin/plugin_fakeMsg"       // 假消息生成插件
	_ "github.com/huoxue1/leafBotPlugin/plugin_friendRequest" // 自动同意好友请求插件
	_ "github.com/huoxue1/leafBotPlugin/plugin_gif"           // gif插件
	_ "github.com/huoxue1/leafBotPlugin/plugin_localStorage"  // 本地存储插件
	_ "github.com/huoxue1/leafBotPlugin/plugin_pixiv"         // 色图查询插件
)

var VERSION = "UnKnow"

func main() {
	go update()
	driver := cqhttp_reverse_ws_driver.NewDriver()
	leafBot.LoadDriver(driver)
	leafBot.InitBots()
	driver.Run()
}

func update() {
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
}
