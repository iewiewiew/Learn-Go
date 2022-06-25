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

/**
Golang 的工作池（Worker Pool）是一种并发设计模式，用于管理和复用一组固定数量的 goroutine（协程），以处理并发任务。工作池可以提高并发执行任务的效率，避免因频繁创建和销毁 goroutine 而产生的开销，同时限制并发任务的数量，以控制系统资源的使用。

工作池通常由以下几个组件组成：
1. 任务队列（Task Queue）：用于存储待处理的任务。任务队列可以是一个有界队列，也可以是一个无界队列，具体取决于应用场景和需求。
1. 工作者（Worker）：代表一个 goroutine，负责从任务队列中获取任务并执行。工作者在启动时会进入一个循环，不断地获取任务并执行，直到任务队列为空或接收到停止信号。
1. 控制器（Controller）：用于管理工作者的数量、任务队列的状态以及整个工作池的生命周期。控制器可以控制工作者的创建和销毁，动态调整工作者的数量，以适应当前的负载情况。

使用工作池的好处包括：
1. 限制并发度：工作池可以限制并发任务的数量，避免系统资源被过度占用，提高系统的稳定性和可靠性。
1. 复用 goroutine：通过复用一组固定数量的 goroutine，可以避免频繁创建和销毁 goroutine 的开销，提高系统的性能和效率。
1. 并发任务调度：工作池可以将任务分配给空闲的工作者，实现并发任务的调度和执行，提高任务处理的并发性。
1. 控制系统负载：通过动态调整工作者的数量，可以根据系统的负载情况来控制并发任务的处理速度，避免过载或资源浪费。

总之，Golang 的工作池是一种有效的并发设计模式，可以提高系统的并发性能和可伸缩性，尤其适用于需要处理大量并发任务的场景。
*/

// 在 Golang 中,job <-chan int 是一个接收(Receive)通道,用于接收类型为 int 的数据.
// job <-chan int 的用法是声明一个只读的通道(Receive-only channel),即它只能用于接收数据,不能用于发送数据.在函数或方法的参数列表中使用 <-chan int 语法可以指定一个只读通道作为参数.
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
	//在 Golang 中,make(chan int, numJobs) 用于创建一个带有缓冲区的通道(Channel).这里的 numJobs 是指定通道的缓冲区大小,表示该通道可以同时容纳的元素数量.
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
