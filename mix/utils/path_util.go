package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/**
 *@author       weimenghua
 *@time         2022-07-30 10:04
 *@description	获取路径
 */

// GetCurrentPath 获取当前路径
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前路径：", dir)
	return dir
}

func GetProjectPath() string {
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("当前项目路径：", pwd)
	return pwd
}

func main() {
	GetCurrentPath()
	GetProjectPath()
}
