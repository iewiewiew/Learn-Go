package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * @author       weimenghua
 * @time         2023-08-24 10:42
 * @description  环境变量
 * 参考资料：https://gobyexample-cn.github.io/environment-variables
 */

// @todo BAR=2 go run go_env.go 没有输出 value
func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
