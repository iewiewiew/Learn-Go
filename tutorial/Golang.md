[TOC]

<h1 align="center">Golang</h1>

> By：weimenghua  
> Date：2022.06.25  
> Description：Go 学习过程

**参考资料**  
- [Go 官网](https://golang.google.cn/doc/)  
- [Go 下载](https://golang.google.cn/dl/)  
- [Go 网址导航](https://hao.studygolang.com/)  
- [Go 系列教程](https://studygolang.com/subject/2)  
- [Go 语言入门教程汇总篇](https://mp.weixin.qq.com/s/ONE4afDx6QAzmdBDKGhdxw)  
- [Go语言进阶，提升必备](https://mp.weixin.qq.com/s/8IDJCE2OOHlic7wsBIfZVA)
- [gobyexample-cn](https://gobyexample-cn.github.io/)  
- [go-by-example](https://www.kancloud.cn/itfanr/go-by-example/81617)
- [LearnGolang](https://github.com/LearnGolang/LearnGolang)  
- [golang-tutorial](https://github.com/nonfu/golang-tutorial)    
- [learngo](https://github.com/inancgumus/learngo)  
- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)  
- [go-ldap-admin](https://gitee.com/eryajf-world/go-ldap-admin)
- [go-recipes](https://github.com/nikolaydubina/go-recipes)



## 1. Go 简介

Go 语言又称 Golang，由 Google 公司于 2009 年发布，近几年伴随着云计算、微服务、分布式的发展而迅速崛起，跻身主流编程语言之列，和 Java 类似，它是一门静态的、强类型的、编译型编程语言，为并发而生，所以天生适用于并发编程（网络编程）。



## 2. Go 环境

### 2.1 环境搭建

**方式一**

```
Mac 环境
brew install go
```

**方式二**

下载 [Go 安装包](https://golang.google.cn/doc/install) 并配置环境变量，本机安装版本为 [go1.18.darwin-arm64](https://golang.google.cn/dl/go1.18.darwin-arm64.pkg)。

```
Mac + Win 环境

mac：/usr/local/go
win：D:\software\Go\bin   

sudo vim ~/.zshrc
export GOROOT=/usr/local/go                     #程序安装的位置
export GOPATH=~/IdeaProjects/Learn-Go           #项目位置
export PATH=$GOROOT/bin:$GOPATH                 #总的路径
source ~/.zshrc

验证Go是否安装成功  
go version

卸载Go
rm -f /usr/local/go
sudo /usr/local/go/uninstall  # 未实践成功

查看Go环境变量
go env
```

**Go 代理**

Go 代理地址：https://goproxy.cn/

```
在终端执行：
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

或者
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
source ~/.profile
```

**镜像地址**

```
七牛镜像
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

阿里镜像
go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```

**初始化项目**

```
go mod init Learn-Go
```

**IDEA 配置 Go环境**

```
1、File > Settings > Plugins ：选择 Go 安装

2、File > Settings > Languages & Frameworks > Go >  GOPATH ：添加 Go 环境变量路径

3、IDEA 添加 Go 模板
File > Settings > Editor > File and Code Templates > Files > Go File
/**
  * @author       weimenghua
  * @time         ${YEAR}-${MONTH}-${DAY} ${TIME}
  * @description
  */
```

**Golang IDE**

[Golang 工具下载](https://www.jetbrains.com/go/nextversion/)

**安装依赖包**

```
安装单个依赖包
go get <package>

使用 `go list` 命令列出项目的所有依赖包，并使用 `go get` 命令批量安装这些依赖包
go list -f '{{ join .Imports "\n" }}' ./... | grep -v "$(go list -f '{{ join .TestImports "\n" }}' ./...)" | grep -v "$(go list -f '{{ join .XTestImports "\n" }}' ./...)" | sort -u | xargs go get -v

go mod download
```

### 2.2 目录结构

- **cmd** 存放主命令文件
- **bin** 存放编译后可执行的文件
- **pkg** 存放编译后的应用包
- **src** 存放应用源代码

在 Go 项目中，cmd 目录通常用于存放主命令文件，也就是应用程序的入口文件。cmd 目录通常包含多个子目录，每个子目录代表一个应用程序，其包含一个或多个 Go 文件，其中包含了应用程序的主要逻辑和入口函数。  
cmd 目录的作用是将应用程序的主要逻辑从其他代码中分离出来，使得应用程序的结构更加清晰和易于维护。在 cmd 目录中，每个子目录都代表一个独立的应用程序，可以在每个子目录中定义自己的命令行标志、参数、环境变量等。

myapp/
├── cmd/
│   ├── myapp/
│   │   ├── main.go
│   │   └── ...
│   └── myotherapp/
│       ├── main.go
│       └── ...
├── internal/
├── pkg/
├── vendor/
├── go.mod
└── go.sum

cmd 目录中的应用程序可以通过 go build 命令进行构建，例如：
```
# 构建 myapp 应用程序
$ go build -o myapp cmd/myapp/main.go

# 构建 myotherapp 应用程序
$ go build -o myotherapp cmd/myotherapp/main.go
```

构建交叉编译的二进制文件，可以使用 -ldflags 和 -tags 选项，例如：
```
# 构建 Linux 平台的二进制文件
$ GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags netgo -o myapp-linux-amd64

# 构建 Windows 平台的二进制文件
$ GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -tags netgo -o myapp-windows-amd64.exe
```



## 3. Go 教程

[Go 知识图谱](https://www.processon.com/mindmap/63f9cf9fe2fac5758f930f43)

### 3.1 语言结构

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

**举例子**

```
package  main

import  "fmt"

func  main()  {
    /* 第一个Go程序 */
    fmt.Println("Hello, World!")
}
```

1. 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。（注：文件内的包名和文件所在包名可以不同）
2. 下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3. 下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4. 下一行 /*...*/ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
5. 下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。使用 fmt.Print("hello, world\n") 可以得到相同的结果。 Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
6. 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

### 3.2 执行程序

运行 go 文件：`go run main.go`  
运行 go 文件：`sudo /usr/local/go/bin/go run main.go` (未配置环境变量需直接指定 go 的绝对路径)  
生成 exe 文件：`go build main.go`  
运行 go 测试文件并生成测试报告和覆盖率报告（cd test）：  
`sudo /usr/local/go/bin/go test -short -v -json -cover -coverprofile cover.out ./... > ../golang-report/report.jsonl`  
`sudo /usr/local/go/bin/go tool cover -html=cover.out -o ../golang-report/index.html`

### 3.3 Makefile

```
.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the application"
	@echo "  run          Run the application"
	@echo "  test         Run the tests"
	@echo "  clean        Clean up the build artifacts"

.PHONY: build
build: ## Build the application
	go build -o myapp cmd/myapp/main.go

.PHONY: run
run: ## Run the application
	go run cmd/myapp/main.go

.PHONY: test
test: ## Run the tests
	go test ./...

.PHONY: clean
clean: ## Clean up the build artifacts
	rm -f myapp
```

在上面的 Makefile 中，我们定义了几个常见的任务，例如：

- `build`：用于构建应用程序，将编译后的二进制文件输出到 `myapp` 文件中。
- `run`：用于运行应用程序，使用 `go run` 命令启动应用程序。
- `test`：用于运行测试，使用 `go test` 命令运行所有的测试。
- `clean`：用于清理构建产物，删除生成的二进制文件。

在 Makefile 中，每个任务都是一个以 `.PHONY` 开头的目标，这个目标定义了任务的名称和任务的描述。任务的具体实现是在下面的命令行中定义的，每个命令行都是以一个制表符（`\t`）开始的，这个制表符是必需的。

使用 Makefile 可以方便地批量执行任务，例如：

```
# 构建应用程序
$ make build

# 运行应用程序
$ make run

# 运行测试
$ make test

# 清理构建产物
$ make clean
```

获取项目版本信息的命令字符串  
```
define get_version
$(shell git describe --tags --always --dirty)
endef
```

### 3.4 依赖管理器

`go mod` 可以用来管理项目中的依赖关系。使用 `go mod` 可以方便地管理和维护项目中使用的第三方库，同时也可以避免由于版本不一致等问题导致的构建失败和运行错误。

`go mod` 的基本操作  
```
初始化项目
$ go mod init [module-name]
其中 [module-name] 是你的项目的模块名称，例如 github.com/yourname/yourproject。使用 go mod init 命令会生成一个 go.mod 文件，用于记录项目依赖关系和版本信息。

添加依赖
$ go get [package-name]
其中 [package-name] 是第三方库的包名，例如 github.com/gin-gonic/gin。使用 go get 命令会自动下载并安装该库，并将其添加到 go.mod 文件中。

升级项目中的依赖库
$ go get -u [package-name]
其中 -u 选项表示升级依赖库的版本。使用 go get -u 命令会自动下载并安装最新版本的依赖库，并将其更新到 go.mod 文件中。

移除依赖
$ go mod tidy
使用 go mod tidy 命令会自动删除项目中未使用的依赖库，并更新 go.mod 文件中的依赖关系。
```

go get 和 go install 的区别

- **`go get`**：主要用于获取和更新依赖包，以及安装依赖包的可执行文件。会修改 `go.mod` 文件。

- **`go install`**：主要用于安装可执行文件，不修改 `go.mod` 文件。

### 3.5 数据库迁移

```
安装数据库迁移库
go get -u github.com/pressly/goose/v3/cmd/goose
go install github.com/pressly/goose/v3/cmd/goose@latest

在迁移目录中创建一个新的迁移文件
goose create create_users_table sql

打开新创建的迁移文件，并编写SQL语句来定义创建用户表的操作。迁移文件分为两类：创建表（up）和删除表（down）。每个迁移文件都需要一个唯一的时间戳前缀，以确保迁移的顺序。
-- <timestamp>_create_users_table.sql

-- +goose Up
-- SQL 在此块中运行时，将应用迁移。

CREATE TABLE users_golang (
   id INT AUTO_INCREMENT PRIMARY KEY,
   username VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL 在此块中运行时，将撤销迁移。

DROP TABLE users_golang;

运行迁移
goose -dir . mysql "youruser:yourpassword@tcp(localhost:3306)/yourdb?parseTime=true" up

回滚迁移
goose -dir . mysql "youruser:yourpassword@tcp(localhost:3306)/yourdb?parseTime=true" down

查看迁移状态
goose -dir migrations mysql "youruser:yourpassword@tcp(localhost:3306)/yourdb?parseTime=true" status
goose -dir migrations mysql "root:123456@tcp(localhost:3306)/dbname?parseTime=true" status
```

### 3.6 godoc 代码文档

```
查看特定包的文档
命令：go doc <package>
例子：go doc fmt

查看特定类型、函数或方法的文档
命令：go doc <package>.<Type/Function/Method>
例子：go doc fmt.Println

安装 godoc
go get golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc@latest

运行 godoc
godoc -http=:6060

查看文档
http://localhost:6060/pkg/
```

### 3.7 swag 文档管理

```
安装依赖
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/iris-contrib/swagger/v12@master
go install github.com/swaggo/swag/cmd/swag@latest

获取模块的路径
go list -f '{{.Dir}}' -m github.com/swaggo/swag

格式化 swag 注解
swag fmt

生成相关的 json 和 yaml 文件
swag init

访问文档
http://127.0.0.1:8088/swagger/index.html#/
```

### 3.8 知识碎片(需调整)

在同一个目录下的 package 名需要统一，例如 utils 目录下的包名统一叫 utils。  
Go 语言中一个程序只能有一个 main 函数，不允许多个 main 函数同时存在。

Go语言本身就不是一个面向对象的编程语言，所以Go语言中没有类的概念，但是他是支持类型的，因此我们可以使用struct类型来提供类似于java中的类的服务，可以定义属性、方法、还能定义构造器。


## 4. Go 框架

常用 [Web 框架](../web/Web.md)：Gin、Beego、Echo、Bali。
