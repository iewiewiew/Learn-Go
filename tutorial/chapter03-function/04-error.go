package main

import (
	"errors"
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2023-02-25 19:42
 * @description  多返回值
 *
 * 参考资料：https://geekr.dev/posts/go-func-params-and-return-values
 */

func add2(a, b *int) (int, error) {
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}
	*a *= 2
	*b *= 3
	return *a + *b, nil
}

func main() {
	x, y := -1, 2
	z, err := add2(&x, &y)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
