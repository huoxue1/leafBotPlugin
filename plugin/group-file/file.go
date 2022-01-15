package group_file

import (
	"fmt"
	"strings"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/fan/utils"
)

func init() {
	utils.CheckDir("./data/temp/")
	plugin := leafbot.NewPlugin("group-file")
	plugin.OnNotice("group_upload", leafbot.Option{
		Weight: 10,
		Block:  false,
	}).Handle(fileHandle)
}

func fileHandle(ctx *leafbot.Context) {
	ctx.SendGroupMsg(ctx.Event.GroupId, []message.MessageSegment{message.Text("让我来看看又上传了什么"), message.Face(289), message.Face(289)})

	// 文件大于100MB
	if ctx.Event.File.Size > 104857600 {
		ctx.SendGroupMsg(ctx.Event.GroupId, []message.MessageSegment{message.Text("文件太大了"), message.Face(314)})
		return
	}

	if strings.HasSuffix(ctx.Event.File.Name, ".zip") {
		err := utils.DownloadFile(ctx.Event.File.FileUrl, "./data/temp/"+ctx.Event.File.Name)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		files, err := utils.GetZipFiles("./data/temp/" + ctx.Event.File.Name)
		if err != nil {
			log.Errorln("获取压缩文件出现了错误" + err.Error())
			return
		}
		ctx.SendGroupMsg(ctx.Event.GroupId, []message.MessageSegment{message.Text(fmt.Sprintf("我看到了这个文件里的所有，包含了：\n") + strings.Join(files, "\r\n"))})
	}
}
