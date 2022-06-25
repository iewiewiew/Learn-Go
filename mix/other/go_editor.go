package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
)

/**
 * @author       weimenghua
 * @time         2024-07-11 15:52
 * @description  编辑器
 */

func main() {
	// 要存储用户输入的变量
	var response string

	// 创建一个编辑器提示
	prompt := &survey.Editor{
		Message:  "Please enter your text:",
		FileName: "temp_editor_file", // 这里你可以指定临时文件的名称
	}

	// 使用编辑器提示获取用户输入
	err := survey.AskOne(prompt, &response)
	if err != nil {
		log.Fatalf("Failed to get user input: %v", err)
	}

	// 输出用户输入
	fmt.Println("You entered:")
	fmt.Println(response)
}
