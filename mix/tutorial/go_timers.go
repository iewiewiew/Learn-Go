package main

import (
	"fmt"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-08-16 11:42
 * @description  定时器
 * 参考资料：https://gobyexample-cn.github.io/timers
 */

// @todo 没看懂
func go_time() {
	// 定时器表示在未来某一时刻的独立事件.你告诉定时器需要等待的时间,然后它将提供一个用于通知的通道. 这里的定时器将等待 2 秒
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C 会一直阻塞,直到定时器的通道 C 明确的发送了定时器失效的值
	<-timer1.C
	fmt.Println("Timer1 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 2 stopepd")
	}
	time.Sleep(2 * time.Second)
}

func main() {
	go_time()
}
