package main

/**
 * @author       weimenghua
 * @time         2023-06-27 14:39
 * @description  gin 例子
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个 Gin 引擎实例
	r := gin.Default()

	// 定义一个 GET 请求的路由
	r.GET("/", func(c *gin.Context) {
		// 使用 c.JSON() 函数将数据结构序列化为 JSON 格式，并返回给客户端
		// gin.H 是 gin 框架中的一个 HTML 模板渲染函数，用于将指定的 HTML 模板渲染并返回给客户端
		//c.JSON(200, gin.H{
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// 启动 Web 服务器并监听 8080 端口
	r.Run(":8080")
}
