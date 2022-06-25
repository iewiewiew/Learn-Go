package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2024-07-15 13:51
 * @description
 */

type tester struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

type StructExample struct {
	Testers []tester `json:"testers"`
	User    tester   `json:"-"` // 当定义一个结构体(struct)时,可以使用 json:"-" 来标记某个字段在编码为 JSON 时应该被忽略。作用：隐藏敏感信息、控制 JSON 输出、兼容性与向后兼容
}

var _internalVar = 42 // 这个变量只能在包内部访问

func myFunction() {
	_temp := 123 // 使用_temp变量,避免与其他变量冲突
	fmt.Println(_temp)
}
