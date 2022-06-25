package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-05 17:54
 * @description  Cobra 构建命令行界面(CLI)工具。示例。
 */

/**
设置了 myapp 命令的别名为 app 和 mycli，用户可以使用这些别名来调用相同的命令。
go build -o go_cmd04 go_cmd04.go
go_cmd04
go_cmd04 myapp
go_cmd04 app
go_cmd04 mycli
*/

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a simple CLI application",
		Long:  "MyApp is a CLI application built with Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			// 处理根命令逻辑
			fmt.Println("Hello from MyApp!")
		},
	}

	rootCmd.Aliases = []string{"app", "mycli"}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
