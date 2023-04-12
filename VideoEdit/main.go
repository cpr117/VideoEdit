// @User CPR
package main

import (
	"VideoEdit/config"
	"VideoEdit/interval"
	"VideoEdit/routes"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {
	interval.InitGlobalVariable()
	g.SetLimit(int(config.Cfg.Server.GLimit))

	// 后台服务
	g.Go(func() error {
		return routes.BackEndServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("all tasks finished successfully")
	}
}
