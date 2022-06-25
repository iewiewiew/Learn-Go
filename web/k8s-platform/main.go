package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"k8s-platform/config"
	_ "k8s-platform/config"
	"k8s-platform/controller"
	"k8s-platform/db"
	"k8s-platform/service"
	_ "k8s-platform/service"
	_ "net/http"
)

/**
 * @author       weimenghua
 * @time         2023-10-27 21:04
 * @description
 */

func main() {
	service.K8s.Init()

	//初始化shujuk
	db.Init()

	//关闭db连接
	defer db.Close()

	//初始化gin对象路由配置
	r := gin.Default()

	//初始化路由规则
	controller.Router.InitApiRouter(r)

	//http server gin程序启动
	r.Run(config.ListenAddr)

}
