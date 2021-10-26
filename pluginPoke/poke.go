package pluginPoke

import (
	"fmt"  //nolint:gci
	"time" //nolint:gci

	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	plugin := leafBot.NewPlugin("戳一戳")
	go plugin.OnNotice("notify").
		AddRule(leafBot.OnlySuperUser).
		SetPluginName("poke").
		AddRule(
			func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) bool {
				if event.SubType != "poke" || event.UserId == event.SelfId || int(event.TargetId) != event.SelfId {
					return false
				}
				return true
			}).SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot leafBot.Api) {
				msg := fmt.Sprintf("服务器使用信息\n\t---------------\nCPU使用率：%0.2f%%\n内存占有率：%.2f%%\n\t----------------", GetCpuPercent(), GetMemPercent())
				if event.GroupId != 0 {
					bot.(leafBot.OneBotApi).SendGroupMsg(event.GroupId, message.Text(msg))
				} else {
					bot.(leafBot.OneBotApi).SendPrivateMsg(event.UserId, message.Text(msg))
				}
			})
}

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() map[string]float64 {
	var content = make(map[string]float64)
	parts, _ := disk.Partitions(true)
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		content[part.Device] = diskInfo.UsedPercent
	}
	return content
}
