// @User CPR
package utils

import (
	"VideoEdit/config"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"path/filepath"
	"strings"
)

type VideoManager struct {
	formatArr []string
}

var (
	VM VideoManager
)

func (vm *VideoManager) formatCheck(target string) bool {
	for _, element := range vm.formatArr {
		if target == element {
			return true
		}
	}
	return false
}

func (vm *VideoManager) NewCutFileName(fileName string) string {
	str := strings.Split(fileName, ".")
	outPath := config.Cfg.Upload.CompletePath
	if ok, _ := FileExists(fmt.Sprintf("%s%s_cut.%s", outPath, str[0], str[1])); !ok {
		return fmt.Sprintf("%s_cut.%s", str[0], str[1])
	} else {
		cnt := 1
		for ok, _ := FileExists(fmt.Sprintf("%s%s_cut(%s).%s", outPath, str[0], string(cnt), str[1])); !ok; {
			cnt += 1
		}
		return fmt.Sprintf("%s_cut(%s).%s", str[0], string(cnt), str[1])
	}
}

// 剪切视频
func (vm *VideoManager) CutVideo(inputVideoPath, startTime, duration string) string {
	outputDir := config.Cfg.Upload.CompletePath
	_, filename := filepath.Split(inputVideoPath)
	tmps := strings.Split(filename, ".")
	forName := tmps[len(tmps)-1]
	if !vm.formatCheck(forName) {
		Logger.Error("Format not supported")
	}
	//uuid, err := uuid.NewUUID()
	//if err != nil {
	//	Logger.Error("generate uuid error:")
	//}

	resultVideoPath := filepath.Join(outputDir, vm.NewCutFileName(filename))
	fmt.Println(resultVideoPath, inputVideoPath)
	err := ffmpeg.Input(inputVideoPath).
		Output(resultVideoPath, ffmpeg.KwArgs{"ss": startTime, "t": duration, "c:v": "copy", "c:a": "copy"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Println(err)
		Logger.Error("VideoCut error:")
	}
	return resultVideoPath
}

func InitVM() {
	VM = VideoManager{}
	VM.formatArr = []string{"mp4", "flv", "mov", "wmv", "mkv", "mpeg", "avi"}
}
