package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

/**
 * @author       weimenghua
 * @time         2023-06-20 10:30
 * @description  http 请求：Get、Post、PostForm、Head 方法
 */

func http_get_demo() {
	resp, err := http.Get("https://xueyuanjun.com")
	if err != nil {
		// 处理错误 ...
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func http_postform_demo() {
	resp, err := http.PostForm("https://xueyuanjun.com/login", url.Values{"name": {"学院君"}, "password": {"test-passwd"}})
	if err != nil {
		// 处理错误
		return
	}

	if resp.StatusCode != http.StatusOK {
		// 处理错误
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func main() {
	//http_get_demo()
	http_postform_demo()
}
