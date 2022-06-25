package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-05 17:23
 * @description  Cobra 构建命令行界面(CLI)工具。标志（flags）是命令行工具中用于传递参数的方式。Cobra 支持多种类型的标志，包括布尔值、字符串、整数等。可以为命令添加标志，并在命令执行时访问这些标志的值。
 */

/**
测试
go run go_cmd03.go info --name Alice --age 30
go run go_cmd03.go info -n Alice -a 30
go run go_cmd03.go info Alice 30 注：name 没有传值
*/
func main() {
	var rootCmd = &cobra.Command{
		Use: "app",
	}

	var name string
	var age int

	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Prints information about a persion",
		//Args:  cobra.ExactArgs(1), // cobra.ExactArgs(1) 是一个用于验证参数数量的函数，它要求命令必须接受恰好一个参数
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Executing infoCmd with argument:", args[0])
			fmt.Printf("Name:%s, Age:%d\n", name, age)
		},
	}

	infoCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the persion")
	infoCmd.Flags().IntVarP(&age, "age", "a", 0, "Age of the persion")

	rootCmd.AddCommand(infoCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
