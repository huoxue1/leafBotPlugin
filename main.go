package main

import (
	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/driver/cqhttp_reverse_ws_driver"

	_ "github.com/huoxue1/fan/plugin/group-mamanger"
	_ "github.com/huoxue1/fan/plugin/sign"
)

func main() {
	driver := cqhttp_reverse_ws_driver.NewDriver()
	leafbot.LoadDriver(driver)
	leafbot.InitBots()
	driver.Run()
}
