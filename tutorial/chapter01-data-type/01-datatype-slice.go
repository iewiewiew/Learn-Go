package main

/**
 * @author       weimenghua
 * @time         2024-07-26 11:18
 * @description  make() 函数用于创建切片、映射和通道。它的语法如下: make(type, len, cap)
 *	type: 需要创建的数据类型,可以是切片、映射或通道。
 *	len: 创建后的初始长度。
 *	cap: 创建后的初始容量(可选)。
 */

import "fmt"

func main() {
	// 1. 创建切片
	// 创建一个长度为5、容量为10的切片
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1) // Output: [0 0 0 0 0]

	// 创建一个长度和容量都为5的切片
	slice2 := make([]string, 5)
	fmt.Println(slice2) // Output: [  ]

	// 创建一个长度为0、容量为10的切片
	slice3 := make([]float64, 0, 10)
	fmt.Println(slice3) // Output: []

	// 2. 创建映射
	// 创建一个初始容量为10的整型到字符串的映射
	m := make(map[int]string, 10)
	m[1] = "one"
	m[2] = "two"
	fmt.Println(m)

	// 3. 创建通道
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
}
