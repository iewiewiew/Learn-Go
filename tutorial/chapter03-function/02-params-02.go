package main

import "fmt"

/**
 * @author      weimenghua
 * @time        2022-06-26 11:48
 * @description	不定长参数
 */

func func1(numbers ...int) { // 把 ... 作用到类型上, 这样就可以约束变长参数的类型
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {
	func1(1, 2, 3, 4, 5)

	slice := []int{1, 2, 3, 4, 5}
	func1(slice...) // 传递切片时需要在末尾加上 ... 作为标识, 表示对应的参数类型是变长参数
	fmt.Println()
	func1(slice[2:5]...) // 输出下标从 2 到 5 的数值
}
