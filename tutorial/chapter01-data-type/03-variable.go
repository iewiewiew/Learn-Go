package main

/**
 * @author        weimenghua
 * @time          2022-06-25 12:40
 * @description	  变量
 * 参考资料: https://geekr.dev/posts/golang-variables-and-constants
 */

import (
	"fmt"
)

func main() {
	var v1 int      // 整型
	var v2 string   // 字符串
	var v3 bool     // 布尔型
	var v4 [10]int  // 数组,数组元素类型为整型
	var v5 struct { // 结构体,成员变量 f 的类型为64位浮点型
		f float64
	}
	var v6 *int            // 指针,指向整型
	var v7 map[string]int  // map(字典),key为字符串类型,value为整型
	var v8 func(a int) int // 函数,参数类型为整型,返回值类型为整型

	// 打印上述变量值
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("v3: %v\n", v3)
	fmt.Printf("v4: %v\n", v4)
	fmt.Printf("v5: %v\n", v5)
	fmt.Printf("v6: %v\n", v6)
	fmt.Printf("v7: %v\n", v7)
	fmt.Printf("v8: %v\n", v8)

	// = 用于变量赋值，使用 = 进行赋值时，左侧的变量必须已经被声明过，否则会引发编译错误。
	// := 它会自动推断变量的类型，并将变量声明和赋值合并为一个语句。
	username := "zhangsan"
	var age int
	age = 18
	fmt.Println(username, age)
}
