package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-02-25 18:19
 * @description  常量
 * 参考资料: https://geekr.dev/posts/golang-variables-and-constants
 */

const ( // iota 被重置为 0
	c0 = iota // c0 = 0
	c1        // c1 = 1
	c2        // c2 = 2
)

const (
	u = iota * 2 // u = 0
	v            // v = 2
	w            // w = 4
)

const x = iota // x = 0
const y = iota // y = 0

const a bool = true
const b int = 1
const c float32 = 1.1
const d float64 = 1.1
const e complex128 = complex(1, 2)
const f string = "字符串"

func main() {
	fmt.Printf("x:%v\n", x)
	fmt.Printf("y:%v\n", y)
	fmt.Println(a, b, c, d, e, f)
}
