package main

import (
	"fmt"
	"math"
)

/**
 * @author       weimenghua
 * @time         2024-07-23 11:34
 * @description
 */

// 定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义 Circle 结构体
type Circle struct {
	Redius float64
}

// 实现 Shape 接口
func (c *Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Redius
}

// 定义 Rectangle 结构体
type Rectangle2 struct {
	Width, Height float64
}

// 实现 Shape 接口
func (r *Rectangle2) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle2) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// 创建 Circle 实例
	c := &Circle{
		Redius: 5,
	}
	r := &Rectangle2{
		Width:  4,
		Height: 5,
	}

	// 使用接口
	shapes := []Shape{c, r}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}
}
