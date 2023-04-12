// @User CPR
package utils

import (
	"VideoEdit/config"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
)

func CreateDir(path string) error {
	fmt.Println(path)
	if ok, err := DirExists(path); !ok {
		log.Println("create %v directory\n" + config.Cfg.Zap.Directory)
		_ = os.Mkdir(config.Cfg.Zap.Directory, os.ModePerm)
	} else {
		log.Println("CreateDir err:", zap.Error(err))
		return err
	}
	return nil
}

// 判断文件目录是否存在
func DirExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	// 存在文件夹
	if fileInfo.IsDir() {
		return true, nil
	}
	return false, nil
}

func FileExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	// 存在文件
	if !fileInfo.IsDir() {
		return true, errors.New("存在同名文件")
	}
	return false, nil
}
