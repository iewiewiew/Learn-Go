package utils

/**
 * @author       weimenghua
 * @time         2023-08-21 15:35
 * @description  时间
 * 参考资料：https://gobyexample-cn.github.io/time
 */

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	secs := now.Unix()
	p(secs)

	nanos := now.UnixNano()
	p(nanos)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))

	then := time.Date(
		2023, 01, 01, 01, 01, 01, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())

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
