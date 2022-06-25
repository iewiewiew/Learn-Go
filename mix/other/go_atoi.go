package main

import (
	"fmt"
	"strconv"
)

/**
 * @author       weimenghua
 * @time         2024-07-09 14:45
 * @description  strconv.Atoi() 是一个将字符串转换为整型的函数。它的全称是 "ASCII to Integer"(ASCII 码转整数)。这个函数最常用于将从输入或配置文件中读取的字符串形式的数字转换为 Go 语言中的整型数据。
 */

func main() {
	i, err := strconv.Atoi("42")
	if err != nil {
		// 处理错误
		fmt.Println(err)
	} else {
		fmt.Println(i) // 输出: 42
	}

	j, err := strconv.Atoi("not a number")
	if err != nil {
		fmt.Println(j, err) // 输出: strconv.Atoi: parsing "not a number": invalid syntax
	}
}
