package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @author       weimenghua
 * @time         2023-08-16 18:20
 * @description  登录接口
 */

// 一条路由规则由三部分组成：http 请求方法、url 路径、控制器函数

// 控制器函数, 通过上下文参数，获取 http 请求参数，响应 http 请求
func doLogin(c *gin.Context) {
	// 获取 post 请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username, password)

	// 通过请求上下文对象 Context, 直接往客户端返回一个字符串
	c.String(200, "username=%s, password=%s", username, password)
}

func main() {
	r := gin.Default()

	// 路由定义 post 请求, url 路径未：/user/login, 绑定 doLogin 控制器函数
	r.POST("/user/login", doLogin)

	r.Run(":8080")
}
