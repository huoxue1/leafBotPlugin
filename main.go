package main

import (
	"fmt"
	"github.com/huoxue1/fan/plugin/model"
	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/driver/cqhttp_positive_ws_driver"

	_ "github.com/huoxue1/fan/plugin/gif"
	"github.com/huoxue1/fan/utils"

	_ "github.com/huoxue1/fan/plugin/flash"
	_ "github.com/huoxue1/fan/plugin/group-mamanger"
	_ "github.com/huoxue1/fan/plugin/sign"
)

func main() {
	driver := cqhttp_positive_ws_driver.NewDriver("ws://127.0.0.1:6700", "qqqq")
	config := utils.GetConfig()
	err := model.InitDb(config.Model.Driver, config.Model.Dsl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	leafbot.LoadDriver(driver)
	leafbot.InitBots(leafbot.Config{
		NickName:     []string{config.Bot.NickName},
		Admin:        0,
		SuperUser:    config.Bot.SuperUser,
		CommandStart: []string{"/"},
		LogLevel:     "info",
	})
	driver.Run()
	select {}
}
