package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * @author       weimenghua
 * @time         2023-08-22 15:59
 * @description  行过滤器
 * 参考资料：https://gobyexample-cn.github.io/line-filters
 */

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println("Input:", ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
