package main

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
}
