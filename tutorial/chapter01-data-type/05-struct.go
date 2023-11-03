package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-12-05 17:50
 * @description  结构体是 Go 语言中对数据进行抽象和封装的主要方法
 */

//结构体基本定义
type Person struct {
	Name    string
	Age     int
	Student Student //嵌套结构体
}

//结构体基本定义 用于嵌套结构体
type Student struct {
	Name  string
	Class string
}

//结构体标签 通常用于配置序列化等操作
type City struct {
	Name string `json:"name"`
	Code int    `json:"code" indo:"yyy"` //多个标签可以叠加
}

/**
可见性规则
结构体中的字段通过首字母大小写控制访问范围:
大写开头的字段在包外可见(public)
小写开头的字段仅包内可见(private)
如果要在包外使用,必须通过提供 getter/setter 方法访问 private 字段:
*/
type Foo struct {
	PublicField  int
	PrivateField int
}

func (foo Foo) GetPrivate() int {
	return foo.PrivateField
}

func (foo Foo) SetPrivate(val int) {
	foo.PrivateField = val
}

//结构体方法
func (p Person) Eat() {
	fmt.Printf(p.Name)
}

func main() {
	//几种创建结构体变量的方式
	p1 := Person{Name: "weimenghua", Age: 18} //命名初始化
	p2 := Person{Name: "weimenghua"}          //部分初始化
	p3 := new(Person)                         //取结构体指针 new()返回对应的结构体指针,需要使用.成员访问符
	var p4 Person                             //声明结构体变量
	fmt.Printf(p1.Name, p2.Name, p3.Name, p4.Name)

	//使用命名字段初始化结构体
	p11 := Person{
		Name: "weimenghua",
		Age:  18,
		Student: Student{
			Name:  "weimenghua",
			Class: "class 1",
		},
	}
	fmt.Printf(p11.Student.Class)

	p22 := Person{}
	p22.Eat()

	//使用零值初始化结构体
	p33 := Person{}
	fmt.Printf(p33.Name)
}
