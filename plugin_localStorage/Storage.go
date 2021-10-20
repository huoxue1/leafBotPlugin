package plugin_localStorage

import (
	"os"

	"github.com/huoxue1/leafBot"
	log "github.com/sirupsen/logrus"
)

func init() {
	go initPath()
}

func initPath() {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("初始检查tmp目录出现无法处理的错误")
			log.Errorln(err)
		}
	}()
	_, err := os.Stat("./tmp")
	if err != nil {
		log.Infoln("正在创建临时文件目录")
		err := os.Mkdir("./tmp", 0666)
		if err != nil {
			log.Errorln("创建临时文件目录失败，将影响某些功能使用")
			return
		}
	}
}

func StorageInit() {
	plugin := leafBot.NewPlugin("local_storage")
	plugin.SetHelp(map[string]string{})
	plugin.OnCommand("storage").
		AddAllies("存储").
		SetBlock(true).
		SetWeight(7).
		AddRule(leafBot.OnlyToMe).
		AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
			if len(state.Args) < 2 {

			}
		})
}
