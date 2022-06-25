package main

/**
 * @author      weimenghua
 * @time        2022-06-25 13:30
 * @description	函数
 */

func calculate1(a int, b int) int { // 两个整型输入, 一个整型输出
	var c = a * b
	return c
}

func calculate2(a, b int) int { // 如果有连续若干个参数, 它们的类型一致, 那么我们无须一一罗列, 只需在最后一个参数后添加该类型.
	var c = a * b
	return c
}

func calculate3(a int, b int) (int, int) { //多返回值
	var c = a * b
	var d = a + b
	return c, d
}

func main() {
	println(calculate1(2, 3))
	println(calculate2(2, 3))
	println(calculate3(2, 3))
}
