package controller

import (
	"fmt"
	"k8s-platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

/**
 * @author       weimenghua
 * @time         2023-11-04 07:53
 * @description
 */

var Pod pod

type pod struct{}

// GetPods (p *pod) 是一个接收器（receiver）的声明（其中 p 是接收器的名称，*pod 是接收器的类型），用于给某个函数指定一个方法。接收器是指定函数的一个参数，它允许函数与特定类型的值进行关联。在这种情况下，(p *pod) 是一个指向类型 pod 的指针的接收器
// Controller中的方法入参是gin.Centext，用于从上下文获取请求参数及定义响应内容
// 流程：绑定参数->调用service代码->根据调用结果响应具体内容
// 获取pod列表，支持过滤、排序、分页
func (p *pod) GetPods(ctx *gin.Context) {
	//声明一个匿名结构体，并实例化了一个该结构体类型的变量params，用于声明入参，get请求为form格式，其他请求为json格式。通过这个匿名结构体的实例化，我们可以在使用这个变量时访问和操作这三个字段。
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	//绑定参数，给匿名结构体中的属性赋值，值是入参
	//form格式(Get)使用ctx.Bind方法，json格式使用ctx.ShouldBindJSON方法
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		//ctx.JSON方法用于返回响应内容，入参是状态码和响应内容，响应内容放入gin.H的map中
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	//service中的方法通过 包名.结构体变量名.方法名 使用，service.Pod.GetPods()
	data, err := service.Pod.GetPods(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Pod列表成功",
		"data": data,
	})
}

// 获取pod详情
func (p *pod) GetPodDetail(ctx *gin.Context) {
	params := new(struct {
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodDetail(params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Pod详情成功",
		"data": data,
	})
}

// 删除pod
func (p *pod) DeletePod(ctx *gin.Context) {
	params := new(struct {
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
	})

	//PUT请求，绑定参数方法改为ctx，ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	fmt.Println("params", params)

	err := service.Pod.DeletedPod(params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "删除pod成功",
		"data": nil,
	})
}

// 更新pod
func (p *pod) UpdatePod(ctx *gin.Context) {
	params := new(struct {
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
	})

	//PUT请求，绑定参数方法改为ctx，ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	fmt.Println("params", params)

	err := service.Pod.UpdatePod(params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "更新pod成功",
		"data": nil,
	})
}

// 获取pod容量
func (p *pod) GetPodContainer(ctx *gin.Context) {
	params := new(struct {
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
	})
	//GET请求，绑定参数改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodContanier(params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取pod容量成功",
		"data": data,
	})
}

// 获取pod中容器日志
func (p *pod) GetPodLog(ctx *gin.Context) {
	params := new(struct {
		ContainerName string `form:"container_name"`
		PodName       string `form:"pod_name"`
		Namespace     string `form:"namespace"`
	})
	//GET请求，绑定参数方法改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodLog(params.ContainerName, params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取pod日志成功",
		"data": data,
	})
}

// 获取每个namespace的pod数量
func (p *pod) GetPodNumPerNp(ctx *gin.Context) {
	data, err := service.Pod.GetPodNumPerNp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取每个namespace的pod数量成功",
		"data": data,
	})
}
