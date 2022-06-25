package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-08-21 11:51
 * @description  通道遍历
 * 参考资料：https://gobyexample-cn.github.io/range-over-channels
 */

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
