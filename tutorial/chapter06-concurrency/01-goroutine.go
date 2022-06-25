package main

/**
 * @author       weimenghua
 * @time         2023-06-05 13:47
 * @description  协程
 *
 * 参考资料：https://gobyexample-cn.github.io/goroutines
 * 使用 go f(s) 在一个协程中调用这个函数. 这个新的 Go 协程将会 并发地 执行这个函数.
 */

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func add(a, b int) {
	var c = a + b
	fmt.Printf("%d + %d = %d", a, b, c)
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	go add(1, 2)
}
