package main

/**
 *@author       weimenghua
 *@time         2022-07-18 17:26
 *@description	读取文件。将整个文件读取到内存。
 */

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("../files/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
