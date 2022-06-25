package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-platform/service"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-11-09 11:18
 * @description
 */

type deployment struct{}

var Deployment deployment

//获取deployment列表，支持过滤、排序、分页
func (d *deployment) GetDeployments(ctx *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
	}
	data, err := service.Deployment.GetDeployments(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Deployment列表成功",
		"data": data,
	})
}

//获取deployment详情
func (d *deployment) GetDeploymentDetail(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `form:"deployment_name"`
		Namespace      string `form:"namespace"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Deployment.GetDeploymentDetail(params.DeploymentName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Deployment详情成功",
		"data": data,
	})
}

//设置deployment副本数
func (d *deployment) ScaleDeployment(ctx *gin.Context) {
	fmt.Printf("ctx", ctx)
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
		ScaleNum       int    `json:"scale_num"`
	})
	//PUT请求，绑定参数方法为ctx.ShouldBindJSON
	if err := ctx.ShouldBind(params); err != nil {
		logger.Error("Bind请求参数，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	fmt.Printf("params", params)
	data, err := service.Deployment.ScaleDeployment(params.DeploymentName, params.Namespace, params.ScaleNum)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "设置Deployment副本数成功",
		"data": fmt.Sprintf("最新副本数：%d", data),
	})
}

//删除deployment
func (d *deployment) DeleteDeployment(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Deployment.DeleteDeployment(params.DeploymentName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "删除Deployment成功",
		"data": nil,
	})
}

//重启deployment
func (d *deployment) RestartDeployment(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
		ImageName      string `json:"image_name"`
		Image          string `json:"image"`
	})
	//PUT请求，绑定参数方法为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Deployment.RestartDeployment(params.DeploymentName, params.Namespace, params.ImageName, params.Image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"mgs":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "重启deployment成功",
		"data": nil,
	})
}

//更新deployment
func (d *deployment) updateDeployment(ctx *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	//PUT方法，绑定参数方法为ctx.ShouldBindJson
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := service.Deployment.UpdateDeployment(params.Namespace, params.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "更新Deployment成功",
		"data": nil,
	})
}
