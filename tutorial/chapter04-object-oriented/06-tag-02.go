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
二、自定义标签创建和处理
1. 格式设计和命名规范
2. 访问和获取标签值
3. 版本标签
*/

//定义一个包含自定义标签的结构体
type Product struct {
	Name     string  `format:"uppercase" validate:"required,min=3,max=100"`
	Price    float64 `format:"currency" validate:"required,min=0"`
	Quantity int     `format:"numeric" validate:"required,min=1"`
}

func processTags(obj interface{}) {
	typ := reflect.TypeOf(obj)

	//遍历结构体字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag

		//输出字段名和自定义标签
		fmt.Printf("Field Name: %s, Custom Tag: %s\n", field.Name, tag.Get("format"))
	}
}

type APIRequest struct {
	Endpoint string `version:"v1"`
	Payload  string `version:"v2"`
}

func getVersion(obj interface{}, version string) string {
	typ := reflect.TypeOf(obj)

	//遍历结构体字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag

		//查找匹配版本的字段
		if tag.Get("version") == version {
			return field.Name
		}
	}

	return "No matching field for the given version."
}

func main() {
	//创建一个包含自定义标签的结构体对象
	product := Product{Name: "Laptop", Price: 1299.99, Quantity: 5}
	//处理自定义标签
	processTags(product)

	//创建一个包含版本标签的结构体对象
	request := APIRequest{Endpoint: "/v1/user", Payload: "data"}

	//获取指定版本的字段名
	fieldName := getVersion(request, "v2")
	fmt.Printf("Field Name for version v2: %s\n", fieldName)
}
