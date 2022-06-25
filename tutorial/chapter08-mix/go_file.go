package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-08-15 18:15
 * @description  文件操作
 */

func go_wirte_file() {
	// 创建文件
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// 写入内容到文件
	_, err = file.WriteString(time.Now().String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created and written successfully.")
}

// @todo 未读取成功
func go_read_file() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建 Scanner 对象
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	line := scanner.Text()
	fmt.Println("文件内容: ", line)

	// 检查是否发生了读取错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
		return
	}

}

func main() {
	//go_wirte_file()
	go_read_file()
}
