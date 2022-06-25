[TOC]

<h1 align="center">Web</h1>

> By：weimenghua  
> Date：2023.06.27  
> Description：Web 框架

**参考资料**  
[Gin 官方文档](https://gin-gonic.com/zh-cn/docs/)   
[gin-vue-admin 开源项目](https://github.com/flipped-aurora/gin-vue-admin/)  



## 1. Gin 框架

Gin 是什么？
Gin 是一个用 Go (Golang) 编写的 HTTP Web 框架。 它具有类似 Martini 的 API，但性能比 Martini 快 40 倍。如果你需要极好的性能，使用 Gin 吧。

初始化项目
```
初始化项目，在 go.mod 首行生成 module example.com/gin_example
go mod init mix.com/mix
go mod init example.com/gin_example

移除依赖
go mod tidy

下载 gin 依赖
go get -u github.com/gin-gonic/gin

下载示例文件
curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
```

gin.Engine 是一个 HTTP web 框架，可以用于构建 Web 应用程序和 HTTP API。它提供了路由、中间件、错误处理、JSON 序列化等功能。

例子
```
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // 定义路由
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, world!",
        })
    })

    // 启动服务
    r.Run(":8080")
}
```

说明：context *gin.Context 是 Go Gin 框架中的一个重要概念，它代表了当前 HTTP 请求的上下文。通过 context *gin.Context，你可以访问和操作当前请求的各种信息，例如请求方法、URL、请求头、请求体、响应状态码等等。

Gin 框架运行模式
为方便调试，Gin 框架在运行的时候默认是 debug 模式，在控制台默认会打印出很多调试日志，上线的时候我们需要关闭 debug 模式，改为 release 模式。
设置Gin框架运行模式：

5.1. 通过环境变量设置
export GIN_MODE=release 
GIN_MODE 环境变量，可以设置为 debug 或者 release

5.2. 通过代码设置
在 main 函数，初始化 gin 框架的时候执行下面代码
// 设置 release 模式
gin.SetMode(gin.ReleaseMode)
// 或者设置 debug 模式
gin.SetMode(gin.DebugMode)

5.3. 常用的 gin.Context 方法

1. c.JSON：发送 JSON 响应。
2. c.IndentedJSON：发送带缩进格式的 JSON 响应。
3. c.String：发送文本响应。
4. c.HTML：发送 HTML 响应。
5. c.Redirect：执行重定向。
6. c.Param：获取 URL 参数。
7. c.Query：获取 URL 查询参数。
8. c.PostForm：获取表单数据。
9. c.Bind：将请求的数据绑定到结构体。
10. c.ShouldBindJSON：将请求的 JSON 数据绑定到结构体。
11. c.Abort：终止请求处理。
12. c.Next：调用下一个处理程序。
13. c.GetHeader：获取请求头信息。
14. c.SetCookie：设置响应的 Cookie。
15. c.Status：设置响应状态码。



## 2. Beego 框架

[Beego 源码](https://github.com/beego/beego)



## 3. Bali 框架

[Bali 框架官网](https://go-bali.github.io/)

`Bali` 是一个基于 Go 语言的 Web 框架，它旨在提供一种简单、易用、高效的方式来构建现代 Web 应用程序和 API。

`Bali` 框架具有以下特点：
- 快速路由：`Bali` 提供了一个快速路由器，支持动态路由和参数解析。
- 中间件支持：`Bali` 提供了中间件机制，可以将多个中间件链接在一起以处理请求。
- RESTful API 开发：`Bali` 对 RESTful API 提供了良好的支持，可以很方便地实现资源的 CRUD 操作。
- 可插拔的模板引擎：`Bali` 支持多种模板引擎，包括 `html/template`、`mustache` 和 `handlebars`。
- 集成 ORM：`Bali` 提供了集成了常见 ORM 框架的支持，例如 `GORM` 和 `XORM`。
- 支持 WebSocket：`Bali` 提供了对 WebSocket 的支持，可以轻松地实现实时通信功能。
- 热加载：`Bali` 支持热加载，可以在不重启服务的情况下更新代码。
- 支持单元测试：`Bali` 提供了一套完整的测试框架，可以方便地编写、运行和管理单元测试。



## 4. 知识碎片

