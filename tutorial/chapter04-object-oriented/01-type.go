package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-02-27 16:20
 * @description  type 定义接口类型.接口类型是一种抽象类型,它定义了一组方法的集合.任何实现了这组方法的具体类型都可以被赋值给该接口类型的变量.
 *
 * type 接口类型名 interface {
 *     方法名1(参数列表1) 返回值列表1
 *     方法名2(参数列表2) 返回值列表2
 *     // ...
 * }
 */

//定义新类型：type关键字可以用来定义新的类型,包括基本类型的别名、结构体、接口等.例如：
type MyInt int

type Person struct {
	Name string
	Age  int
}

type Logger interface {
	Log(message string)
}

//类型定义：在函数签名中,type关键字可以用来定义函数的参数类型.例如：
type BinaryFunc func(int, int) int

func apply(f BinaryFunc, a, b int) int {
	return f(a, b)
}

//其它例子
type Integer int

type Math interface {
	Add(i Integer) Integer
	Mutiply(i Integer) Integer
}

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func (a Integer) Mutiply(b Integer) Integer {
	return a * b
}

func main() {
	var x Integer
	var y Integer
	x, y = 10, 15
	fmt.Println(x.Equal(y))

	var a Integer = 1
	var m Math = &a
	fmt.Println(m.Add(2))
}
