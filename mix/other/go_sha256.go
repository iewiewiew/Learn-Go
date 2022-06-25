package main

import (
	"crypto/sha256"
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2023-08-22 13:58
 * @description  SHA256 散列
 * 参考资料：https://gobyexample-cn.github.io/sha256-hashes
 */

func main() {
	s := "abc"

	h := sha256.New()

	// 写入要处理的字节,如果是一个字符串,需要使用 []byte(s) 将其强制转换成字节数组.
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
