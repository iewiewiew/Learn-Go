package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-06-26 10:45
 * @description  协程
 *
 * 参考资料：https://mp.weixin.qq.com/s/DLofOCOLuYQqoP4mtF1mnw
 */

var counter int = 0

func add1(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add1(1, 2, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched() // 用于手动触发调度器,让出当前 Goroutine 的执行权限,让其他 Goroutine 获取执行机会
		if c > 10 {
			break
		}
	}

	end := time.Now()
	consume := end.Sub(start).Seconds() // end.Sub(start) 是用于计算时间间隔的函数,它返回一个 time.Duration 类型的值,表示 end 和 start 之间的时间间隔.可以通过调用 Seconds() 方法,将时间间隔转换为秒数.
	fmt.Println("程序执行耗时(s)：", consume)
}
