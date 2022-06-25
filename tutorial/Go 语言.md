[TOC]

<h1 align="center">Go 语言</h1>

> By：weimenghua  
> Date：2022.06.25  
> Description：

**参考资料**  
[Go 官网](https://golang.google.cn/doc/)  
[Go 下载](https://golang.google.cn/dl/)  
[Go 网址导航](https://hao.studygolang.com/)  
[Go 系列教程](https://studygolang.com/subject/2)  
[Go 语言入门教程汇总篇](https://mp.weixin.qq.com/s/ONE4afDx6QAzmdBDKGhdxw)   
[golang-tutorial](https://github.com/nonfu/golang-tutorial)  
[learngo](https://github.com/inancgumus/learngo)  
[gobyexample-cn](https://gobyexample-cn.github.io/)  
[LearnGolang](https://github.com/LearnGolang/LearnGolang)    
[go-by-example](https://www.kancloud.cn/itfanr/go-by-example/81617)



## 一、GO 简介
Go 语言又称 Golang，由 Google 公司于 2009 年发布，近几年伴随着云计算、微服务、分布式的发展而迅速崛起，跻身主流编程语言之列，和 Java 类似，它是一门静态的、强类型的、编译型编程语言，为并发而生，所以天生适用于并发编程（网络编程）。



## 二、Go 环境
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

Go代理
https://goproxy.cn/

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

**Golang**
[Golang 工具下载](https://www.jetbrains.com/go/nextversion/)

**安装依赖包**

```
安装单个依赖包
go get <package>

使用 `go list` 命令列出项目的所有依赖包，并使用 `go get` 命令批量安装这些依赖包
go list -f '{{ join .Imports "\n" }}' ./... | grep -v "$(go list -f '{{ join .TestImports "\n" }}' ./...)" | grep -v "$(go list -f '{{ join .XTestImports "\n" }}' ./...)" | sort -u | xargs go get -v
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



## 三、Go 教程

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
1. 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2. 下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3. 下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4. 下一行 /*...*/ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
5. 下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。使用 fmt.Print("hello, world\n") 可以得到相同的结果。 Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
6. 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。



### 3.2 执行程序

运行go文件：`go run main.go`  
运行go文件：`sudo /usr/local/go/bin/go run main.go` (未配置好环境变量直接指定go的绝对路径)  
生成exe文件：`go build main.go`  
运行go测试文件并生成测试报告和覆盖率报告（cd test）：  
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



### 3.5 零散记录(需调整)

### nil

在 Go 语言中，nil 表示一个指针类型的零值，它表示指针不指向任何有效的内存地址。nil 可以用于表示任何指针类型的变量，比如指向一个结构体、数组、切片、函数等的指针。

当一个指针被初始化为 nil，它就没有指向任何有效的内存地址。这意味着它不能被解引用，否则会引发运行时错误。因此，当使用指针时，我们通常需要检查它是否为 nil，以避免潜在的运行时错误。

nil 还可以表示一个接口类型的零值，这意味着该接口没有指向任何具体的实现。这个特性可以用于实现空接口类型，即 interface{}，它可以表示任何类型的值。

总之，nil 在 Go 语言中是一个表示指针类型和接口类型零值的常量。它表示该指针或接口没有指向任何有效的内存地址或具体的实现。

```
var data *int

if condition {
    data = getData()
}

if data != nil {
    // do something with data
}
```

### defer

在 Go 中，defer 语句用于在函数返回时执行一些操作。defer 语句可以被用于释放资源、清理变量、记录日志、统计执行时间等诸多场景。  
注意，当函数中有多个 defer 语句时，它们会按照后进先出（LIFO）的顺序执行。也就是说，最后一个 defer 语句会最先执行，而第一个 defer 语句会最后执行。

```
func main() {
    defer fmt.Println("defer")
    fmt.Println("nomral1")
    fmt.Println("nomral2")
    fmt.Println("nomral3")
}
```

### 指针

在 Go 语言中，& 符号用于取一个对象的地址。如果在一个对象前加上 & 符号，表示获取该对象的地址，并返回一个指向该地址的指针。因此，将 & 用于对象前，可以得到该对象的指针。

在 Go 语言中，指针类型以 * 开头，例如 *int 表示指向 int 类型的指针。因此，如果一个函数的参数类型是 *SomeType，那么它可以接受 SomeType 类型的指针作为参数。

通过将对象的地址传递给函数，可以在函数内部修改对象的值，因为函数拥有指向该对象的指针。这种方式可以用于实现不同的数据结构，如链表和树等。

如何在 Go 中使用指针

```
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println("Before update:", p)

    updatePersonName(&p, "Bob")
    fmt.Println("After update:", p)
}

func updatePersonName(p *Person, name string) {
    p.Name = name
}
```

**Golang 中 = 和 := 的区别**  
= 用于变量赋值，使用 = 进行赋值时，左侧的变量必须已经被声明过，否则会引发编译错误  
:= 它会自动推断变量的类型，并将变量声明和赋值合并为一个语句

在 Go 语言中，struct 的字段标签可以使用反引号 ` 来定义，反引号中的内容被称为标签字符串。标签字符串是在编译时被解析的字符串，可以用于提供额外的元数据，例如字段的名称、验证规则、数据库列名等。

例子

```
type User struct {
    Name     string `json:"name" db:"user_name"`
    Age      int    `json:"age" db:"user_age"`
    Email    string `json:"email" db:"user_email"`
}
```

在 Go 语言中，import 语句用于导入其他包中的代码。有时候可能只需要导入包中的一部分代码，而不是全部代码。在这种情况下，可以使用下划线 _ 替代导入路径中的标识符，以表示忽略该标识符。

```
package main

import (
    _ "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // 使用 _ 来表示导入的包中的标识符不需要被使用
)

func main() {
    // fmt.Println("Hello, world!") // 编译错误：undefined：fmt.Println
}
```


```
file, err := os.OpenFile(repoPath+"/file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
```

- os.O_WRONLY: 以只写模式打开文件。只允许写入文件，不允许读取文件内容。
- os.O_CREATE: 如果文件不存在，则创建文件。如果文件已经存在，则不执行任何操作。
- os.O_TRUNC: 如果文件存在，并且打开模式为写入模式，则截断文件内容为零长度。
- 0644 是文件权限，它指定了新创建的文件的权限。在这个例子中，它表示所有者具有读和写的权限，而其他用户只具有读的权限。


### 工作池

Golang 的工作池（Worker Pool）是一种并发设计模式，用于管理和复用一组固定数量的 goroutine（协程），以处理并发任务。工作池可以提高并发执行任务的效率，避免因频繁创建和销毁 goroutine 而产生的开销，同时限制并发任务的数量，以控制系统资源的使用。

工作池通常由以下几个组件组成：

1. 任务队列（Task Queue）：用于存储待处理的任务。任务队列可以是一个有界队列，也可以是一个无界队列，具体取决于应用场景和需求。

1. 工作者（Worker）：代表一个 goroutine，负责从任务队列中获取任务并执行。工作者在启动时会进入一个循环，不断地获取任务并执行，直到任务队列为空或接收到停止信号。

1. 控制器（Controller）：用于管理工作者的数量、任务队列的状态以及整个工作池的生命周期。控制器可以控制工作者的创建和销毁，动态调整工作者的数量，以适应当前的负载情况。

使用工作池的好处包括：

1. 限制并发度：工作池可以限制并发任务的数量，避免系统资源被过度占用，提高系统的稳定性和可靠性。

1. 复用 goroutine：通过复用一组固定数量的 goroutine，可以避免频繁创建和销毁 goroutine 的开销，提高系统的性能和效率。

1. 并发任务调度：工作池可以将任务分配给空闲的工作者，实现并发任务的调度和执行，提高任务处理的并发性。

1. 控制系统负载：通过动态调整工作者的数量，可以根据系统的负载情况来控制并发任务的处理速度，避免过载或资源浪费。

总之，Golang 的工作池是一种有效的并发设计模式，可以提高系统的并发性能和可伸缩性，尤其适用于需要处理大量并发任务的场景。



在一个名为db的map中查找键为user的值。如果找到了该键，则将对应的值赋给变量value，并将变量ok设置为true。如果没有找到该键，则将变量value设置为该map的值类型的零值，并将变量ok设置为false

```
user := c.Params.ByName("name")
value, ok := db[user]
```



```
安装数据库迁移库
go get -u github.com/pressly/goose/v3/cmd/goose

在迁移目录中创建一个新的迁移文件
goose create create_users_table sql

打开新创建的迁移文件，并编写SQL语句来定义创建用户表的操作
-- <timestamp>_create_users_table.sql

-- +goose Up
CREATE TABLE users (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;

运行数据库迁移命令，将迁移应用到数据库中。在终端中执行以下命令：
goose -dir migrations postgres "user=your_user password=your_password dbname=your_db sslmode=disable" up
```



## 四、Go 框架

常用 [Web 框架](../web/Web 框架.md)：Gin、Echo、Bali。


