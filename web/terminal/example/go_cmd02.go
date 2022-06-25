package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-05 13:50
 * @description  Cobra 构建命令行界面(CLI)工具
 */

func main() {
	rootCmd := &cobra.Command{ // 创建新的 cobra.Command 结构体实例，它代表了应用程序的根命令
		Use:   "myapp",                                                                // 设置命令的名称 myapp
		Short: "A simple command-line application",                                    // 设置命令的简短描述，会在帮助输出中显示
		Long:  `This is a more detailed description of the command-line application.`, // 设置命令的更详细描述，会在帮助输出中显示

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to myapp!")
		},
	}

	// 添加子命令
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Print a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}
	rootCmd.AddCommand(helloCmd)

	// 添加一个带有参数的子命令
	greetCmd := &cobra.Command{
		Use:   "greet [name]",
		Short: "Greet a person by name",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", args[0])
		},
	}
	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	var outputDir = "./docs"

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// 生成文档
	if err := doc.GenMarkdownTree(rootCmd, outputDir); err != nil {
		log.Fatalf("Failed to generate markdown documentation: %v", err)
	}

	log.Println("Markdown documentation generated successfully")

	// 生成文档
	if err := rootCmd.GenBashCompletionFile("go_cmd02.md"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

/**
测试
go run main.go
go run main.go hello
go run main.go greet Alice
*/
