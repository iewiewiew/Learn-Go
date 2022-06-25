package main

import (
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2024-07-23 10:21
 * @description
 */

type People struct {
	Id      int
	Name    string
	Age     int
	Address *Address // 嵌套结构体
}

type Address struct {
	Id   int
	City string
}

func main() {
	// 创建 People 指针类型变量
	s1 := new(People)

	// 创建 People 指针类型变量
	s2 := &People{}

	// 创建 People 类型变量
	s3 := People{}

	// 创建 People 指针类型变量，并根据字段名赋初始值
	s4 := &People{
		Id:   123456,
		Name: "wei",
		Age:  18,
		Address: &Address{
			Id:   123,
			City: "SZ",
		},
	}

	// 创建 People 类型变量，并根据字段在结构体中的顺序赋初始值
	s5 := People{
		123456,
		"wei",
		18,
		&Address{
			Id:   123,
			City: "SZ",
		},
	}

	/** 输出结果：
	s1: &{0  0}
	s2: &{0  0}
	s3: {0  0}
	s4: {123456 wei 18}
	s5: {123456 wei 18}
	*/

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", *s4)
	fmt.Println("s5:", s5)
}
