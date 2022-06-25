package main

/**
 * @author       weimenghua
 * @time         2023-06-27 18:33
 * @description  提交表单
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		// 获取表单参数
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 处理表单参数
		fmt.Println(username, password)

		// 返回响应
		c.JSON(http.StatusOK, gin.H{
			"message": "Form submitted successfully",
		})
	})

	r.Run(":8080")
}

// curl -X POST -H "Content-Type: application/json" -d '{"username":"root", "password":"root"}' http://127.0.0.1:8080/form
