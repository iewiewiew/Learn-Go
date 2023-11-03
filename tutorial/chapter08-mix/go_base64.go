package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

/**
 * @author       weimenghua
 * @time         2023-08-22 14:05
 * @description  Base64 编码
 * 参考资料：https://gobyexample-cn.github.io/base64-encoding
 */

func base64_example() {
	data := "abc"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
}

func UnicodeDecode(raw string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	fmt.Printf(str)
	return str, nil
}

func main() {
	base64_example()
	UnicodeDecode("aa")
}
