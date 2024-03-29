package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-11-04 15:53
 * @description
 */

//处理跨域问题，支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求方法
		method := c.Request.Method
		//添加跨域响应头
		c.Header("Content-Type", "application/json")
		//c.Header("Content-Type", "*")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		//c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Header("Access-Control-Allow-Credentials", "false")
		//旅行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//fmt.Println("c.Request", c.Request)
		//处理请求
		c.Next()
	}
}
