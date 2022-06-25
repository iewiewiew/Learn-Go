package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2024-07-15 14:17
 * @description
 */

/*
在 Go 语言中，nil 表示一个指针类型的零值，它表示指针不指向任何有效的内存地址。nil 可以用于表示任何指针类型的变量，比如指向一个结构体、数组、切片、函数等的指针。
当一个指针被初始化为 nil，它就没有指向任何有效的内存地址。这意味着它不能被解引用，否则会引发运行时错误。因此，当使用指针时，我们通常需要检查它是否为 nil，以避免潜在的运行时错误。
nil 还可以表示一个接口类型的零值，这意味着该接口没有指向任何具体的实现。这个特性可以用于实现空接口类型，即 interface{}，它可以表示任何类型的值。
总之，nil 在 Go 语言中是一个表示指针类型和接口类型零值的常量。它表示该指针或接口没有指向任何有效的内存地址或具体的实现。
*/

func someFunction() (string, error) {
	fmt.Println(123)
	return "true", nil
}

func otherFunction() {
	value, err := someFunction()

	if err != nil {
		fmt.Println("message:", value)
	}
}

func main() {
	otherFunction()
}
