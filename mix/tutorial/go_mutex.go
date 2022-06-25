package main

/**
 * @author       weimenghua
 * @time         2023-08-21 14:21
 * @description  互斥锁：Go语言中的互斥锁(Mutex)是一种用于控制对共享资源的并发访问的机制.它可以确保在任意时刻只有一个 goroutine 可以访问被保护的共享资源,从而避免了数据竞争和并发访问的冲突.
 * 参考资料：https://gobyexample-cn.github.io/mutexes
 */

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 获取互斥锁
			mutex.Lock()

			// 访问共享资源
			counter++

			// 释放互斥锁
			mutex.Unlock()
		}()
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("Counter:", counter)
}
