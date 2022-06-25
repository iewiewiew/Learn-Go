package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-02-25 19:10
 * @description  参数
 *
 * 参考资料：https://geekr.dev/posts/go-func-params-and-return-values
 */

func add(a, b *int) int {
	*a *= 2
	*b *= 3
	return *a + *b
}

func main() {
	x, y := 1, 2
	z := add(&x, &y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
