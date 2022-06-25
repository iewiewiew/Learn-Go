package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-platform/service"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2024-07-20 17:28
 * @description
 */

type example struct {
}

var Example example

func (e *example) GetList(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name"`
		Page  int    `form:"page"`
		Limit int    `form:"limit"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind 请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	data, err := service.Example.GetList(params.Name, params.Page, params.Limit)
	if err != nil {
		logger.Error("获取Example列表失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Example列表成功",
		"data": data,
	})
	logger.Info(data)
}

func (e *example) Create(ctx *gin.Context) {
	var (
		ex  = &service.ExampleCreate{}
		err error
	)

	if err = ctx.ShouldBindJSON(ex); err != nil {
		logger.Error("Bind请求参数ex失败," + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	if err = service.Example.CreateExample(ex); err != nil {
		logger.Error("创建Example失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "创建Example成功",
		"data": nil,
	})
}

func (e *example) DelById(ctx *gin.Context) {
	params := new(struct {
		Id int `json:"id"`
	})

	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}

	if err := service.Example.DelById(params.Id); err != nil {
		logger.Error("删除Example失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "删除Example失败，" + err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "删除Example成功",
		"data": nil,
	})
}
