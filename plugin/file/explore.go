package file

import (
	"fmt"
	"io"

	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	"github.com/huoxue1/fan/utils"
)

type ExploreImpl struct {
}

var (
	client = req.C()
)

func (e *ExploreImpl) Upload(reader io.ReadCloser, fileName string) error {
	config := utils.GetConfig()
	response, err := client.R().SetQueryParams(map[string]string{
		"path":      "/fiction",
		"file_name": fileName,
		"secret":    config.Fiction.UploadSecret,
	}).SetFileReader("file", fileName, reader).
		Post(config.Fiction.RemoteAddress + "/onedrive/upload")
	if err != nil {
		log.Errorln("上传文件出现错误" + err.Error())
		return err
	}
	log.Debugln(response.String())
	return nil
}

func (e *ExploreImpl) Search(word string) []map[string]string {
	config := utils.GetConfig()
	response, err := client.R().
		SetQueryParam("path", "/fiction").
		SetQueryParam("key", word).
		Get(config.Fiction.RemoteAddress + "/onedrive/search")
	if err != nil {
		log.Errorln("获取远程文件错误")
	}
	var files []map[string]string
	fmt.Println(response.String())
	gjson.GetBytes(response.Bytes(), "data").ForEach(func(key, value gjson.Result) bool {

		files = append(files, map[string]string{
			"name": value.Get("name").String(),
			"url":  config.Fiction.RemoteAddress + "/d" + value.Get("path").String(),
		})
		return true
	})
	return files
}
