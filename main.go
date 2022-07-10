package main

import (
	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/driver/cqhttp_default_driver"

	//_ "github.com/huoxue1/fan/plugin/course"
	_ "github.com/huoxue1/fan/plugin/file"
	_ "github.com/huoxue1/fan/plugin/gif"
	"github.com/huoxue1/fan/utils"

	//_ "github.com/huoxue1/fan/plugin/group-file"
	_ "github.com/huoxue1/fan/plugin/group-mamanger"
	_ "github.com/huoxue1/fan/plugin/sign"
)

func main() {
	driver := cqhttp_default_driver.NewDriver()
	config := utils.GetConfig()
	leafbot.LoadDriver(driver)
	leafbot.InitBots(leafbot.Config{
		NickName:     []string{config.Bot.NickName},
		Admin:        0,
		SuperUser:    config.Bot.SuperUser,
		CommandStart: []string{"/"},
		LogLevel:     "info",
	})
	driver.Run()
}
