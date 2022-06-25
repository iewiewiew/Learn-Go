package main

import (
	"fmt"
	"os"
)

/**
 * @author       weimenghua
 * @time         2023-08-23 10:04
 * @description  命令行参数
 * 参考资料：https://gobyexample-cn.github.io/command-line-arguments
 */

// 编译一个可执行二进制文件：go build go_command.go
// 执行程序：./go_command 1 2 3
func main() {
	argsWithProg := os.Args        // os.Args 是一个字符串切片,它包含了程序启动时命令行参数的列表
	argsWithoutProg := os.Args[1:] // 获取除程序路径外的所有命令行参数
	arg := os.Args[3]              // 获取索引 3 命令行参数

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
