package main

import (
	"fmt"
	"reflect"
)

/**
 * @author       weimenghua
 * @time         2023-12-07 09:57
 * @description  结构体标签
 */

/**
一、结构体标签概念
1. 标签定义和格式解析
2. 标签元信息的作用
结构体标签通过在字段上添加特定格式的注释,为字段提供了额外的元信息.在示例中,json 和 db 标签为字段提供了在 JSON 序列化 和 数据库映射 中的配置信息.
3. 常见系统标签样例
Go 语言标准库 和 第三方库中常见的系统标签,如 json、db 等,都遵循一定的格式规范,通过这些标签可以实现一些通用的功能,如序列化、数据库映射等.
*/

//定义一个包含标签的结构体
type User1 struct {
	ID   int    `json:"user_id" db:"id"`
	Name string `json:"user_name" db:"name"`
	Age  int    `json:"user_age" db:"age"`
}

func main() {
	//获取结构体类型
	typ := reflect.TypeOf(User1{})

	//遍历结构体字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag

		//输出字段名、类型和标签信息
		fmt.Printf("Filed Name: %s, Type: %v, JSON tag: %s, DB Tag: %s\n", field.Name, field.Type, tag.Get("json"), tag.Get("db"))
	}
}
