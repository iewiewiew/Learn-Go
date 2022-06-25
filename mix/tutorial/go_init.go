package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2024-07-15 11:33
 * @description  init() 函数
 */

/**
在 Golang 中，func init() 是一种特殊的函数，具有固定的写法和特定的用途。init 函数用于在程序开始执行之前进行初始化操作。以下是一些关于 init 函数的关键点：

自动调用：init 函数会在包被导入时自动调用，不需要手动调用。
无参数和返回值：init 函数不能有参数，也不能有返回值。
每个包可以有多个 init 函数：同一个包内可以有多个 init 函数，甚至在同一个文件中也可以有多个 init 函数。它们的执行顺序是按照它们在文件中出现的顺序。
init 函数的执行顺序：
在单个源文件中，init 函数按它们在文件中出现的顺序执行。
在不同源文件中，init 函数的执行顺序按照文件名的字母顺序执行。
在不同包之间，init 函数的执行顺序是按照包的依赖关系来确定的，即先初始化被依赖的包。
*/

var globalVar int // 声明一个全局变量

// init() 函数用于初始化全局变量
func init() {
	globalVar = 42
	fmt.Println("globalVar initialized to", globalVar)
}

func main() {
	fmt.Println("globalVar value:", globalVar)
}
