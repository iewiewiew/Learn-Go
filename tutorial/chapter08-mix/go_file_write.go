package main

/**
 *@author       weimenghua
 *@time         2022-07-18 17:50
 *@description	写入文件
 */

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("../files/test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("呱呱大王")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
