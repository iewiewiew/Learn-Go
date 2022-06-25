package main

import (
	"errors"
	"fmt"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-16 14:30
 * @description  os.Exit() 函数用于终止当前程序的执行,并返回一个状态码。这个状态码通常被用作程序退出的信号,表示程序是否正常退出。
 */

/**
0: 表示程序正常退出。这是最常见的状态码,表示程序执行成功并按预期退出。
1: 表示程序异常退出。这通常用于表示程序遇到了一些错误或问题,无法正常完成执行。
2: 表示命令行参数有误。当程序无法正确解析命令行参数时,可以返回该状态码。
3-124: 这些状态码通常用于表示特定的错误情况。例如:
3: 表示内存不足
4: 表示系统调用失败
5: 表示I/O错误
等等
128+n: 这些状态码表示程序是由于收到信号而退出的。n 代表信号编号。例如:
128+9: 表示程序收到了 SIGKILL 信号而退出
128+15: 表示程序收到了 SIGTERM 信号而退出
*/

func main() {
	// 模拟一个错误
	err := errors.New("something went wrong")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// 程序正常执行
	fmt.Println("Program executed successfully")
}
