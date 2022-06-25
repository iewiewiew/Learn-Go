package main

/**
 * @author       weimenghua
 * @time         2023-06-20 14:13
 * @description  服务端
 */

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		fmt.Fprintf(writer, "你好, %s", params.Get("name"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动 HTTP 服务器失败: %v", err)
	}
}
