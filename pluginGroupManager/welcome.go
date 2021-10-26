package pluginGroupManager

import "github.com/huoxue1/leafBot"

func init() {
	go WelcomeInit()
}

// WelcomeInit
/**
 * @Description:
 */
func WelcomeInit() {
	plugin.OnNotice("group_increase").SetWeight(10).SetPluginName("入群欢迎").AddHandle(func(event leafBot.Event, bot leafBot.Api) {
		for _, s := range leafBot.DefaultConfig.Plugins.Welcome {
			if s.GroupId == event.GroupId {
				bot.(leafBot.OneBotApi).SendGroupMsg(event.GroupId, s.Message)
			}
		}
	})
}
