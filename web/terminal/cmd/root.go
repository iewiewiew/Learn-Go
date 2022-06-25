package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-04 18:14
 * @description
 */

var RootCmd = &cobra.Command{
	Use:     "啥",
	Short:   "啥啊",
	Long:    "啥啊啊",
	Version: "1.0.0",
}

func init() {
	RootCmd.AddCommand()
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
