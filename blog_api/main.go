package main

import (
	"blog_api/conf"
	"blog_api/db"
	"blog_api/router"
	"blog_api/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	defer db.Db.Close()
	//加载日志
	log := utils.Log()
	gin.SetMode(conf.Conf.Server.Model)

	//路由TODO
	router := router.InitRouter()

	srv := &http.Server{
		Addr:    conf.Conf.Server.Address,
		Handler: router,
	}
	//多线程处理http监听，最大化使用服务器资源
	go func() {
		//启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
		log.Fatalf("listen:%s\n", conf.Conf.Server.Address)
	}()
	quit := make(chan os.Signal)
	//监听📶消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
