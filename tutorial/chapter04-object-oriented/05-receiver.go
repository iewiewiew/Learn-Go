package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-12-06 14:57
 * @description  接收器
 * 方法和接收器是 Go 语言实现面向对象编程的基础。方法将功能定义为特定类型的行为，接收器则是调用方法的对象实例。
 * 接收器可以看作是方法被调用的对象实例。通过接收器，方法与具体的类型实例 associated，从而可以为类型实例添加新的行为。
 */

/**
//其中,ReceiverType 表示接收器的类型,MethodName 为方法的名称。
func (recv ReceiverType) MethodName(params) returnTypes {
  //方法体实现
}
*/

type User struct {
	Id   int
	Name string
}

//接收器的两种形式，值接收器简单直接,而指针接收器可以避免大对象的复制开销并可以在方法内部修改接收器。
//1、接收器为值 User类型上的方法
func (u User) Greet() string {
	return "Hello " + u.Name
}

//2、接收器为指针
func (u *User) Say() string {
	return "Hello " + u.Name
}

func main() {
	u := User{1, "wei"}
	fmt.Printf(u.Greet())
}
