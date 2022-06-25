package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2023-02-27 11:20
 * @description  流程控制语句-选择语句

条件语句：用于条件判断,对应的关键字有 if、else 和 else if;
选择语句：用于分支选择,对应的关键字有 switch、case 和 select;
循环语句：用于循环迭代,对应的关键字有 for 和 range;
跳转语句：用于代码跳转,对应的关键字有 goto.
*/

func switch_demo() {
	tmp := 100
	switch {
	case tmp >= 100:
		fallthrough //想要继续执行后续分支代码,可以通过一个 fallthrough 语句来声明
	case tmp > 90:
		fmt.Printf("tmp > 90")
	case tmp < 90:
		fmt.Printf("tmp < 90")
	default:
		fmt.Printf("我是默认值")
	}
}

func main() {
	switch_demo()
}
