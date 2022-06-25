package main

import (
	"fmt"
	"sort"
)

/**
 * @author       weimenghua
 * @time         2024-07-26 11:30
 * @description  使用 sort 标准库对切片进行排序
 */

type PersionSet struct {
	Name string
	Age  int
}

type ByAge []PersionSet

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []PersionSet{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)
}
