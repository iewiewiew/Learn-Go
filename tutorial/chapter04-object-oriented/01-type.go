package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-02-27 16:20
 * @description  type 定义接口类型。接口类型是一种抽象类型，它定义了一组方法的集合。任何实现了这组方法的具体类型都可以被赋值给该接口类型的变量。

 * type 接口类型名 interface {
 *     方法名1(参数列表1) 返回值列表1
 *     方法名2(参数列表2) 返回值列表2
 *     // ...
 * }
 */

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
