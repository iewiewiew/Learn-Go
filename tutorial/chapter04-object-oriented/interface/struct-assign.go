package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-02-27 16:22
 * @description  面向对象
 *
 * 参考资料：https://geekr.dev/posts/go-type-system
 */

type Integer int

func (a *Integer) Add(b Integer) {
	*a = (*a) + b
}

func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

type Math interface {
	Add(i Integer)
	Multiply(i Integer) Integer
}

func main() {
	var a Integer = 1
	var m Math = &a
	m.Add(2)
	fmt.Printf("1 + 2 = %d\n", a)
}
