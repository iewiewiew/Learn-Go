package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-08-25 11:28
 * @description  使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
 * Go AsciiJSON 是一个用于在 Go 语言中处理 ASCII JSON 数据的库。它提供了将 ASCII JSON 数据解析为 Go 对象以及将 Go 对象转换为 ASCII JSON 数据的功能
 */

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		// map[string]interface{} 是一种用于存储以字符串为键（key）和任意类型为值（value）的映射数据结构
		data := map[string]interface{}{
			"lang": "Go 语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080") // http://127.0.0.1:8080/someJSON
}
