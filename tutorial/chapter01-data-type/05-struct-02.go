package main

/**
 * @author       weimenghua
 * @time         2023-12-11 14:20
 * @description  结构体
 */

// 1. 类型内嵌语法
type A struct {
}

type B A //将类型 A 内嵌到类型 B 中.这样 B 就包含了 A 的所有方法和字段

// 2. 结构体内嵌语法
type Person1 struct {
	Id   int
	Name string
}

type Employee struct {
	Person // 内嵌类型 Person1
	Job    string
}

type Student2 struct {
	Info Info `json:"info"`
}

type Info struct {
	name string `json:"name"`
	age  int    `json:"age"`
}
