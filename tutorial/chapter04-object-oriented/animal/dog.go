package animal

/**
 *@author       weimenghua
 *@time         2022-07-18 12:37
 *@description  Dog 类
 */

type Dog struct {
	*Animal
}

func (d Dog) FavorFood() string {
	return "骨头"
}

func (d Dog) Call() string {
	return "汪汪汪"
}
