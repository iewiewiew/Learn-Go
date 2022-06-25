package main

/**
 * @author       weimenghua
 * @time         2023-02-25 18:33
 * @description	 函数
 *
 * 参考资料：https://geekr.dev/posts/go-func-usage
 * 函数定义：Go 普通函数的基本组成包括：关键字 func、函数名、参数列表、返回值、函数体和返回语句.作为强类型语言,无论是参数还是返回值,在定义函数时,都要声明其类型.
 * new 与 make：new 函数作用于值类型,仅分配内存空间,返回的是指针,make 函数作用于引用类型,除了分配内存空间,还会对对应类型进行初始化,返回的是初始值.在 Go 语言中,引用类型包括切片(slice)、字典(map)和管道(channel),其它都是值类型.
 */

func main() {
	p1 := new(int)    // 返回 int 类型指针,相当于 var p1 *int
	p2 := new(string) // 返回 string 类型指针
	p3 := new([3]int) // 返回数组类型指针,数长度是 3

	type Student struct {
		id    int
		name  string
		grade string
	}
	p4 := new(Student) // 返回对象类型指针

	println("p1: ", p1)
	println("p2: ", p2)
	println("p3: ", p3)
	println("p4: ", p4)

	s1 := make([]int, 3)          // 返回初始化后的切片类型值, 即 []int{0, 0, 0}
	m1 := make(map[string]int, 3) // 返回初始化的字典类型值, 即散列化的 map 结构

	println(len(s1)) // 3
	for i, v := range s1 {
		println(i, v)
	}

	println(len(m1)) // 0
	m1["test"] = 100
	for k, v := range m1 {
		println(k, v)
	}
}
