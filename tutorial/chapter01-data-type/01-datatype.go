package main

/**
 * @author          weimenghua
 * @time            2022-06-25 12:40
 * @description		数据类型: 字符串、整数、数组、字典、切片

map(字典、映射)
map 翻译过来就是字典或者映射, 可以把 map 看做是切片的升级版
只要是可以做 ==、!= 判断的数据类型都可以作为 key (数值类型、字符串、数组、指针、结构体、接口)
map 的 key 的数据类型不能是: slice、map、function
map 和切片一样容量都不是固定的, 当容量不足时底层会自动扩容

切片是用来存储一组相同类型的数据的, map 也是用来存储一组相同类型的数据的
在切片中我们可以通过索引获取对应的元素, 在 map 中我们可以通过 key 获取对应的元素
切片的索引是系统自动生成的, 从 0 开始递增, map 中的 key 需要我们自己指定
map 格式: var dic map[key 数据类型] value 数据类型
*/

import "fmt"

func main() {
	// 1、字符串
	var a string = "张三" // 字符串
	fmt.Println(a)

	// 2、整数
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 3、数组
	var d [8]byte             // 长度为 8 的数组,每个元素为一个字节
	var e [3][3]int           // 二维数组(9 宫格)
	var f [3][3][3]float64    // 三维数组(立体的 9 宫格)
	var g = [3]int{1, 2, 3}   // 声明时初始化
	var h = [...]int{1, 2, 3} // 省略数组长度声明
	var i = new([3]string)    // 通过 new 初始化
	fmt.Println(d, e, f, g, h, i)
	fmt.Println(g[0], g[1], g[2], len(g))
	for i := 0; i < len(g); i++ {
		fmt.Println("第", i, "的值=", g[i])
	}

	// 4、字典
	var j map[int]int = map[int]int{0: 1, 1: 3, 2: 5}
	var k map[string]string = map[string]string{"name": "wei", "age": "18"}
	var l = make(map[string]string) // 表示使用 make 函数创建一个映射 .make 是一个用于创建切片、映射(map)和通道(channel)的内建函数.它用于分配内存并初始化相应的数据结构, 以便能够使用它们进行操作和存储数据.
	l["name"] = "wei"
	l["age"] = "18"
	fmt.Println(j, j[0], j[1], j[2], k, k["name"], l, l["age"])
	for key, value := range l {
		fmt.Println(key, value)
	}

	l["name"] = "wei修改"
	fmt.Println(l["name"])

	delete(l, "name")
	fmt.Println("删除后", l)

	// 5、切片
	// 创建一个整型切片
	var numbers = []int{1, 2, 3, 4, 5}

	// 输出切片的内容
	fmt.Println("切片的内容: ", numbers)

	// 对切片进行追加操作
	numbers = append(numbers, 6, 7, 8)

	// 输出追加后的切片内容
	fmt.Println("追加后的切片内容: ", numbers)
}
