package main

/**
 * @author       weimenghua
 * @time         2023-08-19 08:36
 * @description  测试 HTTP 请求
 */

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func demo() {
	url := "http://127.0.0.1:8080/user/login"

	// 构建请求体
	requestBody := []byte(`{"username":"root","password":"root"}`)

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 输出响应内容 @todo：没有输出参数
	fmt.Println("输出响应内容:", string(body))
}

func main() {
	demo()
}
