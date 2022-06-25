package main

import (
	"encoding/base64"
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2023-08-22 14:05
 * @description  Base64 编码
 * 参考资料：https://gobyexample-cn.github.io/base64-encoding
 */

func main() {
	data := "abc"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
}
