package main

import "fmt"

/**
 * @author      weimenghua
 * @time        2022-06-27 15:08
 * @description 斐波那契
 * 参考资料：https://geekr.dev/posts/go-recursive-function-and-optimization
 */

func fibonacci_demo(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci_demo(n-1) + fibonacci_demo(n-2)
}

func main() {
	n := 10
	num := fibonacci_demo(n)
	fmt.Println(n, num)
}
