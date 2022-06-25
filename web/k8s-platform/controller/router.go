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

// 实例化router结构体，可使用该对象点出首字母大写的方法（包外调用）
var Router router

// 创建router结构体
type router struct{}

// 初始化路由规则
func (r *router) InitApiRouter(router *gin.Engine) {
	//创建测试api接口
	router.GET("/testapi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "testapi success!",
			"data": nil,
		})
	}).
		//登录
		POST("/api/login", Login.Auth).
		GET("/api/k8s/workflows", Workflow.GetList).
		POST("/api/k8s/workflow/create", Workflow.Create).
		DELETE("/api/k8s/workflow/del", Workflow.DelById).
		GET("/api/k8s/examples", Example.GetList).
		POST("/api/k8s/example/create", Example.Create).
		DELETE("/api/k8s/example/del", Example.DelById).
		//namespaces
		GET("/api/k8s/namespaces", Namespace.GetNamespaces).
		GET("/api/k8s/namespaces/detail", Namespace.GetNamespaceDetail).
		//DELETE("api/k8s/namespace/del", Namespace.DeleteNamespace).
		//pod操作
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
		//DELETE("/api/k8s/pod/del", Pod.DeletePod).
		PUT("/api/k8s/pod/update", Pod.UpdatePod).
		GET("/api/k8s/pod/container", Pod.GetPodContainer).
		GET("/api/k8s/pod/log", Pod.GetPodLog).
		GET("/api/k8s/pod/numnp", Pod.GetPodNumPerNp).
		//node操作
		GET("/api/k8s/nodes", Node.GetNodes).
		GET("/api/k8s/node/detail", Node.GetNodeDetail).
		//Ingress操作
		GET("/api/k8s/ingresses", Ingress.GetIngresses)
}
