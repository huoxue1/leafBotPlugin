package pluginFriendRequest

import "github.com/huoxue1/leafBot"

func init() {
	friendRequest()
}

func friendRequest() {
	leafBot.OnRequest("friend").SetWeight(10).SetPluginName("自动同意好友").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		for _, secret := range leafBot.DefaultConfig.Plugins.AutoPassFriendRequest {
			if secret == event.Comment {
				bot.SetFriendAddRequest(event.Flag, true, "")
				return
			}
		}
		bot.SetFriendAddRequest(event.Flag, false, "")
	})
}
