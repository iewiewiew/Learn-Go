package animal

import "fmt"

/**
 *@author       weimenghua
 *@time         2022-07-18 12:36
 *@description  Animal 类
 */

type Animal struct {
	name string
}

//NewAnimal 方法作为 Animal 的构造函数.由于构造函数肯定需要在包以外的地方调用,所以将其首字母大写了,Animal 类包含了 name 属性,首字母小写表示只在 animal 包内可见.
func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.name
}

func main() {
	fmt.Printf(NewAnimal("Tiger").name)
}
