package main

/**
 * @author      weimenghua
 * @time        2022-06-25 12:52
 * @description 计算：加减乘除
 */

import "fmt"

func main() {

	var a int = 1
	var b int = 2
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 1
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}
