package main

import (
	"fmt"
	"net/http"
)

/**
 * @author       weimenghua
 * @time         2023-06-20 14:16
 * @description  客户端
 */

func demo() {
	resp, err := http.NewRequest("GET", "http://127.0.0.1:8080/hello?name=学院君", nil)
	if err != nil {
		// 处理错误 ...
		return
	}

	fmt.Println(resp.Body)

	//defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
}

func main() {
	demo()
}
