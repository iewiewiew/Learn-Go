package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2024-07-15 14:34
 * @description  指针
 */

/**
在 Go 语言中，& 符号用于取一个对象的地址。如果在一个对象前加上 & 符号，表示获取该对象的地址，并返回一个指向该地址的指针。因此，将 & 用于对象前，可以得到该对象的指针。
在 Go 语言中，指针类型以 *（星号） 开头，例如 *int 表示指向 int 类型的指针。因此，如果一个函数的参数类型是 *SomeType，那么它可以接受 SomeType 类型的指针作为参数。
通过将对象的地址传递给函数，可以在函数内部修改对象的值，因为函数拥有指向该对象的指针。这种方式可以用于实现不同的数据结构，如链表和树等。
*/

type Person struct {
	Name string
	Age  int
}

func updatePersonName(p *Person, name string) {
	p.Name = name
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println("Before update:", p)

	updatePersonName(&p, "Bob")
	fmt.Println("After update:", p)
}
