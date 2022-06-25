package main

/**
 * @author      weimenghua
 * @time        2022-06-25 17:44
 * @description	流程控制语句-条件语句

条件语句：用于条件判断,对应的关键字有 if、else 和 else if;
选择语句：用于分支选择,对应的关键字有 switch、case 和 select;
循环语句：用于循环迭代,对应的关键字有 for 和 range;
跳转语句：用于代码跳转,对应的关键字有 goto.
*/

import "fmt"

func condition_demo() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}

	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

func main() {
	condition_demo()
}
