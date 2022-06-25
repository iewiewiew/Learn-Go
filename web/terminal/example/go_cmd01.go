package main

/**
 * @author       weimenghua
 * @time         2024-07-04 17:44
 * @description
 */

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mycli",
		Short: "A simple interactive CLI application",
		Long:  `This is a example interactive CLI application built using the Cobra library in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			runInteractivePrompt()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runInteractivePrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a command (type 'exit' to quit): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			fmt.Println("Exiting...")
			return
		}

		// 在这里添加您的交互式命令逻辑
		fmt.Println("You entered:", command)
	}
}
