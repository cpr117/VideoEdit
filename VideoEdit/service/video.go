// @User CPR
package service

import (
	"VideoEdit/model/req"
	"VideoEdit/utils"
	"VideoEdit/utils/r"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Video struct {
}

func (v *Video) EditVideo(c *gin.Context) (string, int) {
	fmt.Println("EditVideo")
	filepath, videoMsg, code := v.UploadVideo(c)
	if code != r.OK {
		return "", code
	} else {
		resultpath := utils.VM.CutVideo(filepath, videoMsg.Start, videoMsg.Duration)
		return resultpath, r.OK
	}
}

func (*Video) UploadVideo(c *gin.Context) (filepath string, videoMsg req.EditVideo, code int) {
	video := utils.GetFormValue(c, "video")

	utils.Json.Unmarshal(utils.GetFormTest(c, "msg"), &videoMsg)
	utils.Logger.Info("UploadVideo" + videoMsg.Start + videoMsg.Duration)
	filePath, _, uploadErr := utils.Local.UploadFile(video)
	if uploadErr != nil {
		fmt.Println("uploadErr: ", uploadErr)
		return "", videoMsg, r.ERROR_FILE_UPLOAD
	}
	if err := c.SaveUploadedFile(video, filePath); err != nil {
		fmt.Println("err: ", err)
		utils.Logger.Info("UploadVideo", zap.Error(err))
		return "", videoMsg, r.ERROR_FILE_UPLOAD
	}
	//CreateVideo(videoMsg, filePath)
	return filePath, videoMsg, r.OK
}
