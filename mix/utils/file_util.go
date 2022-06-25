package utils

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
 * file, err := os.OpenFile(repoPath+"/file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
 * os.O_WRONLY: 以只写模式打开文件。只允许写入文件，不允许读取文件内容。
 * os.O_CREATE: 如果文件不存在，则创建文件。如果文件已经存在，则不执行任何操作。
 * os.O_TRUNC: 如果文件存在，并且打开模式为写入模式，则截断文件内容为零长度。
 * 0644 是文件权限，它指定了新创建的文件的权限。在这个例子中，它表示所有者具有读和写的权限，而其他用户只具有读的权限。
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
