package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2023-06-20 18:05
 * @description  json 序列化
 */

type User struct {
	Name    string
	Website string
	Age     uint
	Male    bool
	Skills  []string
}

func json_demo1() {
	user := User{
		"学院君",
		"https://xueyuanjun.com",
		18,
		true,
		[]string{"Golang", "PHP", "C", "Java", "Python"},
	}

	u, err := json.Marshal(user)

	if err != nil {
		fmt.Println("JSON 编码失败：%v\n", err)
		return
	}

	fmt.Println("JSON 格式数据：%s\n", string(u)) // string(u)：输出JSON字符串，而不是数字
}

func json_demo2() {
	user := User{
		"学院君",
		"https://xueyuanjun.com",
		18,
		true,
		[]string{"Golang", "PHP", "C", "Java", "Python"},
	}
	u, err := json.Marshal(user)

	var user2 User
	err = json.Unmarshal(u, &user2) // & 符号用于取一个对象的地址
	if err != nil {
		fmt.Println("JSON 解码失败：%v\n", err)
	}

	fmt.Println("JSON 解码结果：%#v\n", user2)
}

func main() {
	json_demo1()
	json_demo2()
}
