package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-10-27 23:31
 * @description
 */

//实例化router结构体，可使用该对象点出首字母大写的方法（包外调用）
var Router router

//创建router结构体
type router struct{}

//初始化路由规则
func (r *router) InitApiRouter(router *gin.Engine) {
	//创建测试api接口
	router.GET("/testapi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "testapi success!",
			"data": nil,
		})
	}).
		GET("/api/k8s/workflows", Workflow.GetList).
		GET("/api/k8s/namespaces", Namespace.GetNamespaces).
		GET("/api/k8s/namespaces/detail", Namespace.GetNamespaceDetail).
		DELETE("api/k8s/namespace/del", Namespace.DeleteNamespace)
}
