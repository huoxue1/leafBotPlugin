package pluginfriendRequest

import "github.com/huoxue1/leafBot"

func init() {
	go friendRequest()
}

func friendRequest() {
	plugin := leafBot.NewPlugin("好友申请")
	plugin.SetHelp(map[string]string{"": "自动同意好友请求"})

	plugin.OnRequest("friend").SetWeight(10).SetPluginName("自动同意好友").AddHandle(func(event leafBot.Event, bot leafBot.Api) {
		for _, secret := range leafBot.DefaultConfig.Plugins.AutoPassFriendRequest {
			if secret == event.Comment {
				bot.(leafBot.OneBotApi).SetFriendAddRequest(event.Flag, true, "")
				return
			}
		}
		if len(leafBot.DefaultConfig.Plugins.AutoPassFriendRequest) == 0 {
			bot.(leafBot.OneBotApi).SetFriendAddRequest(event.Flag, true, "")
			return
		}
		bot.(leafBot.OneBotApi).SetFriendAddRequest(event.Flag, false, "")
	})
}
