// @User CPR
package utils

import (
	"VideoEdit/config"
	"go.uber.org/zap"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

type LocalUtil struct{}

var Local = new(LocalUtil)

func (*LocalUtil) UploadFile(file *multipart.FileHeader) (filePath, fileName string, err error) {
	filePath = config.Cfg.Upload.StorePath
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		Logger.Info("文件夹不存在，创建文件夹")
		fileErr := os.MkdirAll(filePath, os.ModePerm)
		if fileErr != nil {
			Logger.Error("创建文件夹失败, err:", zap.Error(fileErr))
			return "", "", err
		}
	}
	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	fileName = fileId + path.Ext(file.Filename)
	filePath = filePath + fileName
	return filePath, fileName, nil
}
