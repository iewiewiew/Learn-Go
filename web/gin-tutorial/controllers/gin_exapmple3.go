package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2024-07-15 13:43
 * @description
 */

func main() {
	r := gin.Default()

	// 模拟数据库
	db := map[string]string{
		"alice": "Hello Alice!",
		"bob":   "Hello, Bob!",
	}

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user] // 在一个名为db的map中查找键为user的值。如果找到了该键，则将对应的值赋给变量value，并将变量ok设置为true。如果没有找到该键，则将变量value设置为该map的值类型的零值，并将变量ok设置为false。
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User no found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": value,
		})
	})
	r.Run()
}
