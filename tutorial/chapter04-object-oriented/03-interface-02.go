package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-12-11 13:43
 * @description  接口
 */

//1、接口声明格式
//声明一个接口
type Stringer interface {
	String() string
}

//自定义类型实现接口
type myString string

func (m myString) String() string {
	return string(m)
}

//2、方法集合并
type A interface {
	a()
}

type B interface {
	b()
}

type C interface {
	A
	B
}

type Test struct {
}

func (t Test) a() {}
func (t Test) b() {}

func main() {
	var i Stringer

	i = myString("Hello")

	fmt.Println(i.String())

	//////////////////////////

	var c C = Test{}
	c.a()
	c.b()

	fmt.Println(c) //打印出C
}
