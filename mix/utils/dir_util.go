package utils

import (
	"fmt"
	"os"
)

/**
 * @author       weimenghua
 * @time         2023-08-22 16:18
 * @description  目录操作
 * 参考资料：https://gobyexample-cn.github.io/directories
 */

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	// defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file.txt")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	c, err := os.ReadDir("./")
	for _, entry := range c {
		fmt.Println("", entry.Name(), entry.IsDir())
	}
}
