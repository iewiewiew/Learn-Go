package main

/**
 * @author       weimenghua
 * @time         2023-06-25 16:49
 * @description  recover：在 Go 语言中,recover() 函数用于从 panic 中恢复.当函数调用了 panic() 函数时,程序会中断执行,并且开始沿着调用栈向上查找能够处理该 panic 的 defer 函数,如果找到了能够处理 panic 的 defer 函数,那么该函数就会被执行,否则整个程序就会退出.
 */

import (
	"fmt"
)

func main() {
	foo()
	fmt.Println("Program execution continues...")
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting foo...")
	bar()
	fmt.Println("Finishing foo...")
}

func bar() {
	fmt.Println("Starting bar...")
	panic("Something went wrong!")
	fmt.Println("Finishing bar...")
}
