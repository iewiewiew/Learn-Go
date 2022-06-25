package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-12-05 17:50
 * @description  结构体是 Go 语言中对数据进行抽象和封装的主要方法（Golang 结构体 ～ 其它语言的类）
 */

/**
在 Go 语言中，结构体（struct）是一种用户自定义的复合数据类型，用于封装一组相关的字段。结构体可以包含零个或多个字段，每个字段可以是不同的数据类型。结构体的定义形式如下：
```go
type Person struct {
    Name string
    Age  int
}
```

理解 Go 结构体的关键概念包括：
1. **定义结构体**：使用`type`关键字来定义一个新的结构体类型。
2. **结构体字段**：结构体可以包含零个或多个字段，每个字段有一个字段名和对应的数据类型。
3. **实例化结构体**：通过结构体类型可以创建结构体的实例（也称为结构体变量）。例如：`var p Person`。
4. **访问结构体字段**：使用`.`操作符来访问结构体实例的字段，例如：`p.Name`。
5. **匿名结构体**：可以定义没有字段名的结构体，用于临时的数据封装。

结构体在 Go 语言中被广泛应用，用于表示复杂的数据结构，如数据库记录、HTTP请求/响应等。结构体提供了一种组织和封装数据的方式，有助于编写清晰、可维护的代码。
*/

// Person 结构体基本定义
type Person struct {
	Name    string
	Age     int
	Student Student //嵌套结构体
}

// Student 结构体基本定义 用于嵌套结构体
type Student struct {
	Name  string
	Class string
}

// City 结构体标签 通常用于配置序列化等操作。在 Go 语言中，struct 的字段标签可以使用反引号 ` 来定义，反引号中的内容被称为标签字符串。标签字符串是在编译时被解析的字符串，可以用于提供额外的元数据，例如字段的名称、验证规则、数据库列名等。
type City struct {
	Name string `json:"name"`
	Code int    `json:"code" indo:"yyy"` //多个标签可以叠加
}

// Foo 可见性规则 结构体中的字段通过首字母大小写控制访问范围:大写开头的字段在包外可见(public)、小写开头的字段仅包内可见(private)、如果要在包外使用, 必须通过提供 getter/setter 方法访问 private 字段
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

// Eat 结构体方法
func (p Person) Eat() {
	fmt.Printf(p.Name)
}

func main() {
	// 几种创建结构体变量的方式
	p1 := Person{Name: "wei", Age: 18} // 命名初始化
	p2 := Person{Name: "wei"}          // 部分初始化
	p3 := new(Person)                  // 取结构体指针 new() 返回对应的结构体指针, 需要使用.成员访问符
	var p4 Person                      // 声明结构体变量
	fmt.Printf("p1: %s, p2: %s, p3: %s, p4: %s\n", p1.Name, p2.Name, p3.Name, p4.Name)

	// 使用命名字段初始化结构体
	p11 := Person{
		Name: "wei",
		Age:  18,
		Student: Student{
			Name:  "wei",
			Class: "class 1",
		},
	}
	fmt.Printf("p11: ", p11.Student.Class+"\n")

	p22 := Person{}
	p22.Eat()

	// 使用零值初始化结构体
	p33 := Person{}
	fmt.Printf("p33: ", p33.Name)

	// 匿名结构体
	// 定义并初始化匿名结构体指针
	params := new(struct { // new(struct { ... }) 使用 new 函数分配内存。new 函数分配内存并返回一个指向该匿名结构体的指针。返回的指针类型为 *struct { ... }，指向一个零值初始化的结构体实例。
		Namespace string `form:"namespace"`
		Page      int    `form:"page"`
		Limit     int    `form:"limit"`
	})

	// 设置结构体字段的值
	params.Namespace = "example"
	params.Page = 1
	params.Limit = 10

	fmt.Println("\n")
	fmt.Println(params.Namespace)
	fmt.Println(params.Page)
	fmt.Println(params.Limit)
}
