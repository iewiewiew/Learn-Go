package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12/httptest"
	"github.com/wonderivan/logger"
	"k8s-platform/service"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2024-03-07 10:29
 * @description
 */

var Ingress ingress

type ingress struct{}

//获取ingress列表，支持过滤、分页、排序
func (i *ingress) GetIngresses(ctx *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace""`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(httptest.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	data, err := service.Ingress.GetIngresses(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(httptest.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Ingress列表成功",
		"data": data,
	})
}
