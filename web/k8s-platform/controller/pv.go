package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	service "k8s-platform/service"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-12-05 16:07
 * @description
 */

type Pv pv
type pv struct{}

func (p *pv) GetPvs(ctx *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Page      int    `json:"page"`
		Limit     int    `json:"limit"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "请求参数错误",
			"data": nil,
		})
	}
	data, err := service.Pv.GetPvs(params.Namespace, params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取pv列表失败",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取pv列表成功",
		"data": data,
	})
}

func (p *pv) GetPvDetail(ctx *gin.Context) {
	params := new(struct {
		pvName string `json:"pv_name"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败！", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pv.GetPvDetail(params.pvName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取pv详情成功",
		"data": data,
	})
}

//删除pv
func (p *pv) DeletePv(ctx *gin.Context) {
	params := new(struct {
		pvName string `json:"pv_name"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Pv.DeletePv(params.pvName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "删除pv成功",
		"data": nil,
	})
}
