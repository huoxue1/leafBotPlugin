package plugin_course

import (
	"os"
	"time"

	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	log "github.com/sirupsen/logrus"
)

func init() {
	cmd()
}

func cmd() {
	plugin := leafBot.NewPlugin("课程表")
	plugin.OnCommand("课表", leafBot.Option{
		PluginName: "课程表",
		Weight:     10,
		Block:      true,
		Allies:     []string{"课程表"},
		Rules:      nil,
	}).AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		_, err := os.Stat("./config/course.yml")
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		course, err := getCourse(time.Now())
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		event.Send(message.Image("base64://" + draw(course)))
	})
}
