package main

import "github.com/gin-gonic/gin"

/**
 * @author       weimenghua
 * @time         2023-06-28 15:36
 * @description
 */

type LoginForm struct {
	User     string `form:"User" binding:"required"`
	Password string `from:"User" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var form LoginForm

		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run(":8080")
}

// curl -X POST -H "Content-Type: application/json" -d '{"user":"user", "password":"password"}' http://127.0.0.1:8080/login
