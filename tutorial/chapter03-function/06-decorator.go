package main

import (
	"fmt"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-02-27 11:32
 * @description  装饰器
 *
 * 参考资料：https://geekr.dev/posts/go-decorator-implement-by-high-order-function
 */

// MultiPlyFunc 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int

// 乘法运算函数
func multiply(a, b int) int {
	return a * b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()
		c := f(a, b) // 执行乘法运算函数
		end := time.Since(start)
		fmt.Printf("--- 执行耗时：%v ---\n", end)
		return c
	}
}

func main() {
	a := 2
	b := 8
	c := multiply(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)

	// 通过装饰器调用乘法函数,返回的是一个匿名函数
	decorator := execTime(multiply)
	// 执行修饰器返回函数
	c1 := decorator(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c1)
}
