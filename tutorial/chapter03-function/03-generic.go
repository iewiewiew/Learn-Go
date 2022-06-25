package main

import (
	"fmt"
	"reflect"
)

/**
 * @author       weimenghua
 * @time         2023-02-25 19:31
 * @description  泛型 generic
 *
 * 参考资料：https://geekr.dev/posts/go-func-params-and-return-values
 */

func generic_demo(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Println("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	generic_demo(1, "1", [1]int{1}, true)
}
