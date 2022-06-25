package main

import (
	_ "example/docs"
	"github.com/gin-gonic/gin"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/**
 * @author       weimenghua
 * @time         2024-07-20 12:49
 * @description
 */

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8088")
}
