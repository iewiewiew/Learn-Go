package main

import (
	"fmt"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-08-16 17:07
 * @description  工作池
 * 参考资料：https://gobyexample-cn.github.io/worker-pools
 */

//@todo 没太懂

//在 Golang 中，job <-chan int 是一个接收（Receive）通道，用于接收类型为 int 的数据。
//job <-chan int 的用法是声明一个只读的通道（Receive-only channel），即它只能用于接收数据，不能用于发送数据。在函数或方法的参数列表中使用 <-chan int 语法可以指定一个只读通道作为参数。
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	//在 Golang 中，make(chan int, numJobs) 用于创建一个带有缓冲区的通道（Channel）。这里的 numJobs 是指定通道的缓冲区大小，表示该通道可以同时容纳的元素数量。
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
