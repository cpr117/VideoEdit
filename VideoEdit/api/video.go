// @User CPR
package api

import (
	"VideoEdit/utils/r"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type Video struct{}

var upgrader = websocket.Upgrader{}

func (*Video) EditVideo(c *gin.Context) {
	retpath, code := videoService.EditVideo(c)
	r.ReturnFile(c, code, "", retpath)
}

func (*Video) echo(c *gin.Context) {
	//服务升级，对于来到的http连接进行服务升级，升级到ws
	cn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer cn.Close()
	if err != nil {
		panic(err)
	}
	for {
		mt, message, err := cn.ReadMessage()
		if err != nil {
			log.Println("server read:", err)
			break
		}
		log.Printf("server recv msg: %s", message)
		msg := string(message)
		if msg == "woshi client1" {
			message = []byte("client1 去服务端了一趟")

		} else if msg == "woshi client2" {
			message = []byte("client2 去服务端了一趟")
		}

		err = cn.WriteMessage(mt, message)
		if err != nil {
			log.Println(" server write err:", err)
			break
		}
	}
}
