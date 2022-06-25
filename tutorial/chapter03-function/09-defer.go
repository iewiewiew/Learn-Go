package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-06-20 11:46
 * @description  defer

 * 在 Go 中,defer 语句用于在函数返回时执行一些操作.defer 语句可以被用于释放资源、清理变量、记录日志、统计执行时间等诸多场景.
 * 注意,当函数中有多个 defer 语句时,它们会按照后进先出(LIFO)的顺序执行.也就是说,最后一个 defer 语句会最先执行,而第一个 defer 语句会最后执行.
 */

func main() {
	defer fmt.Println("defer")
	fmt.Println("nomral1")
	fmt.Println("nomral2")
	fmt.Println("nomral3")
}
