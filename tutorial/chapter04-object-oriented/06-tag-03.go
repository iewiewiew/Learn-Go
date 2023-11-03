package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2023-12-08 16:54
 * @description  结构体标签
 */

/**
三、标签应用场景
1. 控制 JSON 序列化
2. ORM 映射数据库
3. 表格数据关联
*/

type Persion struct {
	Name string `json:"persion_name"`
	Age  int    `json:"persion_age"`
}

type Employee struct {
	ID   int    `db:employee_id`
	Name string `db:employee_name`
	Role string `db:employee_role`
}

type ProductInfo struct {
	Name     string `table:"product_name"`
	Price    string `table:"product_price"`
	Quantity string `table:"product_quantity"`
}

func main() {
	//创建一个包含JSON标签的结构体对象
	persion := Persion{
		Name: "Alice",
		Age:  25,
	}
	//输出JSON格式字符串
	jsonData, _ := json.Marshal(persion)
	fmt.Println(string(jsonData))

	//创建一个包含数据库映射标签的结构体对象
	//employee := Employee{ID:101, Name:"Bob", Role:"Engineer"}
	//构造数据库查询语句
	//query := constructQuery(employee)
	//fmt.Println("Database Query:", query)

	//创建一个包含表格数据关联标签的结构体对象
	//productInfo := ProductInfo{Name: "Laptop", Price: "$1299.99", Quantity: "5"}
	//输出表格格式数据
	//tableData := generateTable(productInfo)
	//fmt.Println(tableData)
}
