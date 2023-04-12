// @User CPR
package test

import (
	req2 "VideoEdit/model/req"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"
)

var (
	ctx = context.Background()
)

func TestRegister(t *testing.T) {
	r := gin.Default()

	// 设置路由
	r.GET("/", func(c *gin.Context) {
		// 创建请求数据
		//reg := req.Register{
		//	UserName: "cpr123123123",
		//	Password: "321321321",
		//	Email:    "cpr19990622@gmail.com",
		//	Code:     "",
		//	Phone:    "13676511152",
		//}
		//
		//// 将数据编码为 JSON 格式
		////jsonStr, err := json.Marshal(reg)
		//if err != nil {
		//	c.AbortWithError(http.StatusInternalServerError, err)
		//	return
		//}
		c.JSON(200, gin.H{
			"UserName": "cpr123123123",
			"Password": "321321321",
			"Email":    "cpr19990622@gmail.com",
			"Code":     "",
			"Phone":    "13676511152",
		})
	})
	r.Run(":8080")
}

func TestSendCode(t *testing.T) {
	//r := gin.Default()
	//r.GET("localhost:8765/sendCode", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"username": "cpr123123123",
	//		"email":    "cpr19990622@gmail.com",
	//	})
	//})
	//r.Run(":8080")

	router := gin.Default()

	// 定义一个路由处理函数，发送包含JSON数据的HTTP请求
	router.GET("/send-json", func(c *gin.Context) {
		// 创建一个HTTP客户端
		client := &http.Client{}

		// 创建一个包含JSON数据的结构体
		reqcode := req2.RegisterCode{
			UserName: "cpr123123123",
			Email:    "cpr19990622@gmail.com",
		}

		// 将结构体转换为JSON格式
		requestBody, err := json.Marshal(reqcode)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error marshaling JSON: %s", err.Error()))
			return
		}

		// 创建一个HTTP请求
		req, err := http.NewRequest("POST", "127.0.0.1:8765/sendCode", bytes.NewBuffer(requestBody))
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating request: %s", err.Error()))
			return
		}

		// 设置请求头
		req.Header.Set("Content-Type", "application/json")

		// 发送HTTP请求
		resp, err := client.Do(req)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error sending request: %s", err.Error()))
			return
		}

		defer resp.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return
		}

		// 将响应体返回给客户端
		c.String(http.StatusOK, string(body))
	})

	// 启动Gin服务器
	router.Run(":8080")
}

func TestTest(t *testing.T) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8090", Path: "/echo"}
	log.Printf("client1 connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial server:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		reqcode := req2.RegisterCode{
			UserName: "cpr123123123",
			Email:    "cpr19990622@gmail.com",
		}
		c.WriteJSON(reqcode)
		//for {
		//	_, message, err := c.ReadMessage()
		//	if err != nil {
		//		log.Println("client read err:", err)
		//		return
		//	}
		//	log.Printf("client recv msg: %s", message)
		//}
	}()

	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	for {
		select {
		// if the goroutine is done , all are out
		case <-done:
			return
		case <-time.Tick(time.Second * 5):
			err := c.WriteMessage(websocket.TextMessage, []byte("woshi client2"))
			if err != nil {
				log.Println("client write:", err)
				return
			}
		case <-interrupt:
			log.Println("client interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("client1 write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
