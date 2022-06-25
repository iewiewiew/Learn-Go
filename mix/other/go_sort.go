package main

import (
	"fmt"
	"sort"
)

/**
 * @author       weimenghua
 * @time         2023-08-21 14:45
 * @description  排序
 * 参考资料：https://gobyexample-cn.github.io/sorting
 */

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{3, 2, 1}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
