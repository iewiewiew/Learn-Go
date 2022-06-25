package main

/**
 * @author       weimenghua
 * @time         2023-11-26 09:01
 * @description
 */

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type InferenceOutput struct {
	Data []string `json:"data"`
}

//定义了一个函数类型 inferResultFunc.这个函数类型接受两个参数,第一个参数是指向 gin.Context 的指针,第二个参数是 InferenceOutput 结构体的切片.这个函数类型可以用来表示一个接受上下文和输出数据的函数,这样可以在代码中更方便地使用和传递这个函数类型.
type inferResultFunc func(c *gin.Context, outputs []InferenceOutput)

func main() {
	// 创建一个示例的输出数据
	output := InferenceOutput{
		Data: []string{"result1", "result2", "result3"},
	}

	// 创建一个示例的处理函数
	handlerFunc := func(c *gin.Context, outputs []InferenceOutput) {
		// 在这里可以编写处理输出数据的逻辑
		for _, output := range outputs {
			dataJson, _ := json.Marshal(output)
			c.String(200, string(dataJson))
		}
	}

	// 调用处理函数
	router := gin.Default()
	router.GET("/inference", func(c *gin.Context) {
		handlerFunc(c, []InferenceOutput{output})
	})
	router.Run(":8080")
}
