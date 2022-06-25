package main

import (
	"fmt"
	"sync"
)

/**
 * @author       weimenghua
 * @time         2023-08-21 13:54
 * @description  WaitGroup：等待多个协程完成
 * 参考资料：https://gobyexample-cn.github.io/waitgroups
 */

func main() {
	var wg sync.WaitGroup

	// 启动三个 goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d 执行完毕\n", id)
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("所有 goroutine 执行完毕")
}
