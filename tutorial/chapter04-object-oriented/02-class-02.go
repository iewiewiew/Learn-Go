package main

/**
 * @author      weimenghua
 * @time        2022-06-25 16:02
 * @description	结构体

 * 定义结构体类型
 * type 类型名称 struct{
 *   属性名称 属性类型
 *   属性名称 属性类型
 *  ... ...
 * }
 */

import "fmt"

func main() {
	type Student struct {
		name string
		age  int
	}

	type Demo struct {
		age int               //基本类型作为属性
		arr [3]int            //数组类型作为属性
		sce []int             //切片类型作为属性
		mp  map[string]string //字典类型作为属性
		stu Student           //结构体类型作为属性
	}

	var d Demo = Demo{
		33,
		[3]int{1, 3, 5},
		[]int{2, 4, 6},
		map[string]string{"class": "one"},
		Student{
			"lnj",
			33,
		},
	}
	fmt.Println(d)

	type Demo1 struct {
		name string
		age  int
	}
	var tmp = Demo1{}
	fmt.Println(tmp)

	//完全初始化
	var demo1 = Demo1{"wei", 18}
	fmt.Println(demo1)
	fmt.Println(demo1.name, demo1.age)

	//部分初始化
	var demo2 = Demo1{name: "wei1"}
	fmt.Println(demo2)

	var demo3 = struct {
		name string
		age  int
	}{
		name: "wei",
		age:  18,
	}
	fmt.Println(demo3)
}
