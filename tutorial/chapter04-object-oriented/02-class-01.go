package main

import "fmt"

/**
 * @author      weimenghua
 * @time        2022-06-27 15:16
 * @description 面向对象
 */

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

// NewStudent 通过定义形如 NewXXX 这样的全局函数（首字母大写）作为类的初始化函数
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

func main() {
	student := NewStudent(1, "wei", false, 100)
	fmt.Println(student)
}
