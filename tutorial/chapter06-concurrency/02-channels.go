package main

/**
 * @author       weimenghua
 * @time         2023-06-05 14:51
 * @description  通道
 *
 * 参考资料：https://gobyexample-cn.github.io/channels
 * 通道(channels) 是连接多个协程的管道. 你可以从一个协程将值发送到通道,然后在另一个协程中接收.
 */

import (
	"fmt"
	_ "fmt"
)

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
