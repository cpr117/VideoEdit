// @User CPR
package routes

import (
	"VideoEdit/config"
	"log"
	"net/http"
	"time"
)

// 后台服务
func BackEndServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
