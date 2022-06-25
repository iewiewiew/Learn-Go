package main

/**
 * @author      weimenghua
 * @time        2022-06-25 12:52
 * @description 流程控制语句-循环语句

条件语句：用于条件判断,对应的关键字有 if、else 和 else if;
选择语句：用于分支选择,对应的关键字有 switch、case 和 select;
循环语句：用于循环迭代,对应的关键字有 for 和 range;
跳转语句：用于代码跳转,对应的关键字有 goto.
*/

import "fmt"

func for_expample() {
	arr := [...]int{1, 3, 5}

	//传统for循环遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	//for...range循环遍历
	for i, v := range arr {
		fmt.Println(i, v)
	}

	//for true {
	//	fmt.Print("这是无线循环 \n")
	//}
}

func for_underline() {
	numbers := []int{1, 2, 3, 4, 5}

	//在 Go 语言中,for 循环中的下划线 _ 可以用作占位符,用于表示忽略循环中的某个变量.
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func for_underline2() (int, int, int) {
	return 1, 2, 3
}

func main() {
	//for_expample()
	for_underline()

	a, _, c := for_underline2() // 使用 _ 忽略了第二个整数的值
	fmt.Println(a, c)
}
