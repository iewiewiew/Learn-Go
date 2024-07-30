package main

/**
 * @author      weimenghua
 * @time        2022-06-25 14:27
 * @description 九九乘法表
 */

import "fmt"

func ninenie_demo() {
	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 { // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}

	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2) //位宽为8,左对齐
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	ninenie_demo()
}
