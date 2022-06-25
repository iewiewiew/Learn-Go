package main

import (
	"fmt"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-08-15 11:58
 * @description	 获取时间
 */

func main() {
	// 获取当前时间
	currentTime := time.Now()

	// 格式化并打印当前时间
	fmt.Println("当前时间:", currentTime)

	// 获取年份、月份、日等信息
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()

	fmt.Printf("年份: %d\n", year)
	fmt.Printf("月份: %s\n", month)
	fmt.Printf("日期: %d\n", day)
	fmt.Printf("小时: %d\n", hour)
	fmt.Printf("分钟: %d\n", minute)
	fmt.Printf("秒数: %d\n", second)
}
