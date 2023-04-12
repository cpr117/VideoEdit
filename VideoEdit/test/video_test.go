// @User CPR
package test

import (
	"VideoEdit/utils"
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestVideoCut(t *testing.T) {
	testlist := []string{"mp4", "flv", "mov", "wmv", "mkv", "avi"}
	for _, format := range testlist {
		filename := fmt.Sprintf("./video/upload/%s_test.%s", format, format)
		_result := utils.VM.CutVideo(filename, "0:10", "5")
		fmt.Println(_result)
	}
}

func TestHttpVideoCut(t *testing.T) {
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1dWlkIjoiM2YzZThlYzgyZWVkODBiNzRjMzI3NjIwYmM0YmExMDYiLCJpc3MiOiJWaWRlb0VkaXQiLCJleHAiOjE2ODEzNzA3Mzh9.C9OEK8YizwL-JZxS-VarPe5F0OmL9TIMgXnGrJo7e24"
	url := "http://127.0.0.1:8765/api/editvideo"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("../video/prepare/avi_test.avi")
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("video", filepath.Base("../video/prepare/avi_test.avi"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("msg", "{\"start\":\"10\",\"duration\":\"10\"}")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	outfile, err := os.Create("../video/test_complete/avi_test.avi")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestSocketVideoCut(t *testing.T) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8765/videoeditsocket", nil)
	if err != nil {
		fmt.Println("Failed to establish WebSocket connection:", err)
		return
	}
	defer conn.Close()

	// 打开文件
	file, err := os.Open("../video/prepare/avi_test.avi")
	if err != nil {
		utils.Logger.Error("文件打开失败：", zap.Error(err))
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	// 发送请求相关信息
	err = conn.WriteMessage(websocket.TextMessage, utils.Json.MarshalByte(map[string]string{
		"filename": fileInfo.Name(),
		"start":    "5",
		"duration": "10",
	}))
	if err != nil {
		utils.Logger.Error("发送数据失败：", zap.Error(err))
	}

	video, _ := ioutil.ReadAll(file)

	err = conn.WriteMessage(websocket.BinaryMessage, video)
	if err != nil {
		utils.Logger.Error("发送文件失败：", zap.Error(err))
	}

	//file, err := os.Open("../video/prepare/avi_test.avi")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//
	//fileInfo, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//fileSize := fileInfo.Size()
	//
	//buffer := make([]byte, fileSize)
	//_, err = file.Read(buffer)
	//if err != nil {
	//	panic(err)
	//}
	//err = conn.WriteMessage(websocket.BinaryMessage, buffer)
	//if err != nil {
	//	panic(err)
	//}
	//
	//messageType, p, err := conn.ReadMessage()
	//if err != nil {
	//	panic(err)
	//}
}
