package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-12-06 14:57
 * @description  接收器
 */

/**
在 Go 语言（Golang）中，接收器（Receiver）是用于定义方法（Method）的特殊参数。它可以让你将方法与特定类型关联起来，从而实现类似面向对象编程中的类和方法的功能。接收器有两种类型：值接收器和值指针接收器。

// ReceiverType 表示接收器的类型，MethodName 为方法的名称
func (recv ReceiverType) MethodName(params) returnTypes {
  //方法体实现
}
*/

type User struct {
	Id   int
	Name string
}

// Greet 接收器的两种形式，值接收器简单直接，而指针接收器可以避免大对象的复制开销并可以在方法内部修改接收器
// 1、接收器为值 User 类型上的方法。值接收器：方法接收的是接收器的一个副本，修改副本不会影响原始值。
func (u User) Greet() string {
	return "Hello " + u.Name
}

// Say 2、接收器为指针。指针接收器：方法接收的是接收器的指针，修改指针所指向的值会影响原始值。
func (u *User) Say() string {
	return "Hello " + u.Name
}

type Rectangle struct {
	Width, Height int
}

// Area 定义一个方法，使用值接收器。定义了一个 Rectangle 结构体，并为它定义了一个 Area 方法。这个方法使用了值接收器 r Rectangle，所以在方法内部 r 是 rect 的一个副本，任何对 r 的修改都不会影响到 rect。
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Scale 定义一个方法，使用指针接收器。定义了一个 Scale 方法，它使用了指针接收器 r *Rectangle。因为 r 是一个指向 rect 的指针，所以在方法内部对 r 的修改会直接影响到 rect。
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	u := User{1, "wei"}
	fmt.Println(u.Greet())
	fmt.Println(u.Say())

	// 创建一个 Rectangle 实例
	rect := Rectangle{3, 4}
	fmt.Println("Area: ", rect.Area())

	// 调用 Scale 方法
	rect.Scale(2)

	// 输出缩放后的尺寸
	fmt.Println("Width:", rect.Width)
	fmt.Println("Height:", rect.Height)
}
