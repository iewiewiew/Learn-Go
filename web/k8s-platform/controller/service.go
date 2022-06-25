package controller

import (
	"k8s-platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

/**
 * @author       weimenghua
 * @time         2023-11-30 10:45
 * @description
 */

var Service1 service1

type service1 struct{}

//获取service列表，支持分页、排序、过滤
func (s *service1) GetServices(ctx *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespaces string `form:"namespace"`
		Page       int    `form:"Page"`
		Limit      int    `form:"limit"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Service.GetServices(params.FilterName, params.Namespaces, params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Service列表成功",
		"data": data,
	})
}

//获取Service详情
func (s *service1) GetServiceDetail(ctx *gin.Context) {
	params := new(struct {
		ServiceName string `form:"service_name"`
		Namespace   string `form:"namespace"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	data, err := service.Service.GetServiceDeatil(params.ServiceName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取service详情成功",
		"data": data,
	})
}

//创建service
func (s *service1) GetServicedetail(ctx *gin.Context) {
	var (
		serviceCreate = new(service.ServiceCreate)
		err           error
	)
	if err = ctx.ShouldBindJSON(serviceCreate); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	if err = service.Service.CreateService(serviceCreate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "创建Service成功",
		"data": nil,
	})
}

//删除service
func (s *service1) DeleteService(ctx *gin.Context) {
	params := new(struct {
		ServiceName string `json:"service_name"`
		Namespace   string `json:"namespace"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Service.DeleteService(params.ServiceName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "删除service成功",
		"data": nil,
	})
}

//更新service
func (s *service1) UpdateService(ctx *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败！")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Service.UpdateService(params.Content, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "更新Service成功",
		"data": nil,
	})
}
