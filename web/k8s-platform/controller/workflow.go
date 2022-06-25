package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-platform/service"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-10-28 14:53
 * @description
 */

type workflow struct{}

var Workflow workflow

// 获取列表分页查询
func (w *workflow) GetList(ctx *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
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

	data, err := service.Workflow.GetList(params.FilterName, params.Namespace, params.Page, params.Limit)
	if err != nil {
		logger.Error("获取Workflow列表失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Workflow列表成功",
		"data": data,
	})
	logger.Info(data)
}

// 创建workflow
func (w *workflow) Create(ctx *gin.Context) {
	var (
		wc  = &service.WorkflowCreate{}
		err error
	)

	if err := ctx.ShouldBindJSON(wc); err != nil {
		logger.Error("Bind 请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	if err = service.Workflow.CreateWorkflow(wc); err != nil {
		logger.Error("创建Workflow失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": "创建Workflow成功",
		})
	}
}

func (w *workflow) DelById(ctx *gin.Context) {
	params := new(struct {
		Id int `json:"id"`
	})
	if err := ctx.ShouldBind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	if err := service.Workflow.DelById(params.Id); err != nil {
		logger.Error("删除workflow失败，", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "删除workdlow成功",
		"data": nil,
	})
}
