package utils

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func CheckDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, 0666)
		if err != nil {
			return
		}
	}
}

func DownloadFile(url, file string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	open, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
		_ = open.Close()
	}()
	_, err = io.Copy(open, response.Body)
	if err != nil {
		return err
	}
	return err
}

func GetZipFiles(zipFile string) ([]string, error) {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()
	var files []string
	for _, f := range zipReader.File {
		if f.Flags == 0 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			files = append(files, string(content))
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			files = append(files, f.Name)
		}

	}
	return files, err
}
