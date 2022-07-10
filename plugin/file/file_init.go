package file

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	"github.com/huoxue1/fan/plugin/model"
	"github.com/huoxue1/fan/utils"
)

func init() {
	plugin := leafbot.NewPlugin("文件管理")
	utils.CheckDir("./data/temp/")

	plugin.OnRequest("friend").Handle(func(ctx *leafbot.Context) {
		ctx.SetFriendAddRequest(ctx.Event.Flag, true, "")
	})

	plugin.OnFullMatchGroup("文件初始化", leafbot.Option{
		Weight: 0,
		Block:  true,
		Allies: nil,
		Rules:  []leafbot.Rule{leafbot.OnlySuperUser},
	}).Handle(fileInit)

	plugin.OnStart("找书").Handle(search)
	plugin.OnNotice("group_upload", leafbot.Option{
		Weight: 10,
		Block:  false,
	}).Handle(fileHandle)
}

type Explore interface {
	Upload(reader io.ReadCloser, filename string) error
	Search(word string) []map[string]string
}

type File struct {
	GroupId       int    `json:"group_id"`
	FileId        string `json:"file_id"`
	FileName      string `json:"file_name"`
	Busid         int    `json:"busid"`
	FileSize      int    `json:"file_size"`
	UploadTime    int    `json:"upload_time"`
	DeadTime      int    `json:"dead_time"`
	ModifyTime    int    `json:"modify_time"`
	DownloadTimes int    `json:"download_times"`
	Uploader      int64  `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
}

type Folder struct {
	GroupId        int    `json:"group_id"`
	FolderId       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int    `json:"create_time"`
	Creator        int64  `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int    `json:"total_file_count"`
}

func search(ctx *leafbot.Context) {
	utils.CheckDir("./temp/")
	defer func() {
		_ = os.RemoveAll("./temp/")
	}()
	explore := &ExploreImpl{}
	fictionName := strings.TrimPrefix(ctx.Event.Message.ExtractPlainText(), "找书")
	files := explore.Search(fictionName)
	fmt.Println(files)
	if len(files) == 0 {
		ctx.Send(message.Text("未能找到该类型的书"))
	}
	i := 0

	eventChan, closer := ctx.GetMoreEvent(func(ctx1 *leafbot.Context) bool {
		if ctx1.GroupID == ctx.GroupID && ctx1.UserID == ctx.UserID {
			return true
		}
		return false
	})
	defer closer()
	for {
		msg := fmt.Sprintf("第%d-%d条，共%d条\n", i+1, i+10, len(files))
		if len(files) > i+10 {
			for j := i; j < i+10; j++ {
				msg += fmt.Sprintf("%d : %v\n", j+1, files[j]["name"])
			}
			msg += "\n发送获取[编号,编号即可获取小说],例如获取10,20\n\n"
			msg += "获取小说请先添加机器人为好友！"
		} else if len(files) <= i+10 {
			for j := i; j < len(files); j++ {
				msg += fmt.Sprintf("%d : %v\n", j+1, files[j]["name"])
			}
			msg += "\n发送获取[编号,编号即可获取小说],例如获取10,20\n\n"
			msg += "获取小说请先添加机器人为好友！\n\n"
			msg += "已经是最后一页了额！"
		}
		msg_id := ctx.Send(message.ReplyWithMessage(int64(ctx.Event.MessageID), message.Text(msg)))
		select {
		case data := <-eventChan:
			{

				if data.Message.ExtractPlainText() == "下一页" {
					i += 10
					ctx.DeleteMsg(msg_id)
					continue
				} else if strings.HasPrefix(data.Message.ExtractPlainText(), "获取") {
					gets := strings.Split(strings.TrimPrefix(data.Message.ExtractPlainText(), "获取"), ",")
					fmt.Println(gets)
					ctx.DeleteMsg(msg_id)
					s := new(model.Sign)
					s.QQ = ctx.UserID
					_ = model.Query(s)
					if int(s.Fraction) < len(gets)*2 {
						ctx.Send([]message.MessageSegment{message.Text("积分不足！建议通过签到获取！"), message.At(ctx.UserID)})
						return
					}
					model.UpdateFraction(ctx.UserID, int64(-(len(gets) * 2)))
					_ = model.Query(s)
					ctx.Send(message.Text(fmt.Sprintf("正在为你下载小说，已扣除积分%d，剩余积分%d", 2*len(gets), s.Fraction)))
					for _, get := range gets {
						index, _ := strconv.Atoi(get)
						if index < 0 || index >= len(files) {
							return
						} else {
							url := files[index]["url"]
							err := utils.DownloadFile(url, "./temp/"+files[index]["name"])
							if err != nil {
								continue
							}
							dir, _ := os.Getwd()
							ctx.UploadPrivateFile(ctx.UserID, fmt.Sprintf("%v/temp/%v", dir, files[index]["name"]), files[index]["name"])
						}
					}
					return
				} else {
					return
				}
			}
		case <-time.After(time.Minute):
			{
				ctx.DeleteMsg(msg_id)
				return
			}
		}

	}
}

func fileHandle(ctx *leafbot.Context) {
	explore := &ExploreImpl{}
	ctx.SendGroupMsg(ctx.Event.GroupId, []message.MessageSegment{message.Text("让我来看看又上传了什么"), message.Face(289), message.Face(289)})

	// 文件大于200MB
	if ctx.Event.File.Size > 204857600 {
		ctx.SendGroupMsg(ctx.Event.GroupId, []message.MessageSegment{message.Text("文件太大了"), message.Face(314)})
		return
	}

	if strings.HasSuffix(ctx.Event.File.Name, ".zip") {
		// 下载压缩包文件
		err := utils.DownloadFile(ctx.Event.File.FileUrl, "./data/temp/"+ctx.Event.File.Name)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		// 退出时删除文件
		defer os.Remove("./data/temp/" + ctx.Event.File.Name)

		// 获取压缩包中的所有文件
		files, err := utils.GetZipFiles("./data/temp/" + ctx.Event.File.Name)
		if err != nil {
			log.Errorln("获取压缩文件出现了错误" + err.Error())
			return
		}
		ctx.SendGroupMsg(ctx.Event.GroupId, []message.MessageSegment{message.Text(fmt.Sprintf("我看到了这个文件里的所有，包含了：\n") + strings.Join(files, "\r\n"))})
		// 将压缩包文件上传
		reader, err := zip.OpenReader("./data/temp/" + ctx.Event.File.Name)
		if err != nil {
			log.Errorln("打开压缩文件失败" + err.Error())
			return
		}
		for _, file := range reader.File {
			if !strings.HasSuffix(file.Name, ".txt") {
				continue
			}
			if file.NonUTF8 {
				log.Warningln("文件未知编码")
				continue
			}
			fileData, err := file.Open()
			if err != nil {
				log.Errorln("打开压缩包内部文件失败" + err.Error())
				continue
			}
			log.Infoln("开始上传" + file.Name)
			err = explore.Upload(fileData, file.Name)
			if err != nil {
				log.Errorln("文件上传失败" + err.Error())
			}
			_ = fileData.Close()
		}
		err = reader.Close()
		if err != nil {
			log.Errorln("关闭压缩文件句柄失败" + err.Error())
			return
		}
		err = os.Remove("./data/temp/" + ctx.Event.File.Name)
		if err != nil {
			log.Errorln("本地缓存压缩文件删除失败" + err.Error())
		}
	} else if strings.HasSuffix(ctx.Event.File.Name, ".txt") {
		resp, err := http.Get(ctx.Event.File.FileUrl)
		if err != nil {
			log.Errorln(fmt.Sprintf("文件” %v“下载失败", ctx.Event.File.Name))
			return
		}
		err = explore.Upload(resp.Body, ctx.Event.File.Name)
		if err != nil {
			log.Errorln(fmt.Sprintf("文件上传失败%v", ctx.Event.File.Name))
		}
		log.Infoln(fmt.Sprintf("文件%v上传成功", ctx.Event.File.Name))
	}
}

func fileInit(ctx *leafbot.Context) {
	explore := &ExploreImpl{}
	datas := map[string]*File{}
	zipDatas := map[string]*File{}
	files := ctx.GetGroupRootFiles(ctx.GroupID)
	files.Get("files").ForEach(func(key, value gjson.Result) bool {
		f := new(File)
		_ = json.Unmarshal([]byte(value.String()), f)
		if strings.HasSuffix(f.FileName, ".txt") {
			datas[value.Get("file_id").String()] = f
		} else if strings.HasSuffix(f.FileName, ".zip") {
			zipDatas[value.Get("file_id").String()] = f
		}
		return true
	})
	files.Get("folders").ForEach(func(key, value gjson.Result) bool {
		filesByFolder := ctx.GetGroupFilesByFolder(ctx.GroupID, value.Get("folder_id").String())
		filesByFolder.Get("files").ForEach(func(key, value gjson.Result) bool {
			f := new(File)
			_ = json.Unmarshal([]byte(value.String()), f)
			if strings.HasSuffix(f.FileName, ".txt") {
				datas[value.Get("file_id").String()] = f
			} else if strings.HasSuffix(f.FileName, ".zip") {
				zipDatas[value.Get("file_id").String()] = f
			}
			return true
		})
		return true
	})
	handleZip(ctx, zipDatas, explore)
	handleTxt(ctx, datas, explore)
}

func handleZip(ctx *leafbot.Context, datas map[string]*File, explore Explore) {
	utils.CheckDir("./temp/")
	for _, f := range datas {

		log.Infoln(f.FileName)
		url := ctx.GetGroupFileUrl(ctx.GroupID, f.FileId, f.Busid).Get("url").String()
		if url == "" {
			continue
		}
		log.Infoln(url)
		err := utils.DownloadFile(url, "./temp/"+f.FileName)
		if err != nil {
			log.Errorln("文件下载失败" + err.Error())
			continue
		}
		reader, err := zip.OpenReader("./temp/" + f.FileName)
		if err != nil {
			log.Errorln("打开压缩文件失败" + err.Error())
			continue
		}
		for _, file := range reader.File {
			if !strings.HasSuffix(file.Name, ".txt") {
				continue
			}
			if file.NonUTF8 {
				log.Warningln("文件未知编码")
				continue
			}
			fileData, err := file.Open()
			if err != nil {
				log.Errorln("打开压缩包内部文件失败" + err.Error())
				continue
			}
			log.Infoln("开始上传" + file.Name)
			err = explore.Upload(fileData, file.Name)
			if err != nil {
				log.Errorln("文件上传失败" + err.Error())
			}
			_ = fileData.Close()
		}
		err = reader.Close()
		if err != nil {
			log.Errorln("关闭压缩文件句柄失败" + err.Error())
			continue
		}
		err = os.Remove("./temp/" + f.FileName)
		if err != nil {
			log.Errorln("本地缓存压缩文件删除失败" + err.Error())
		}

	}

}

func handleTxt(ctx *leafbot.Context, datas map[string]*File, explore Explore) {
	//explore := &ExploreImpl{}
	for _, f := range datas {
		log.Infoln(f.FileName)
		url := ctx.GetGroupFileUrl(ctx.GroupID, f.FileId, f.Busid).Get("url").String()
		if url == "" {
			continue
		}
		log.Infoln(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Errorln(fmt.Sprintf("文件” %v“下载失败", f.FileName))
			continue
		}
		err = explore.Upload(resp.Body, f.FileName)
		if err != nil {
			log.Errorln(fmt.Sprintf("文件上传失败%v", f.FileName))
		}
		log.Infoln(fmt.Sprintf("文件%v上传成功", f.FileName))

	}
}
