package main

import (
	"fmt"
	"net/url"
)

/**
 * @author       weimenghua
 * @time         2023-08-21 19:20
 * @description  URL 解析
 * 参考资料：https://gobyexample-cn.github.io/url-parsing
 */

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Host)
	fmt.Println(u.Port())
	fmt.Println(u.Path)
}
