package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @author       weimenghua
 * @time         2024-07-16 14:57
 * @description  捕获 Ctrl + C 信号
 */

func main() {
	// 捕获 Ctrl + C 信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	fmt.Println("Starting...")

	// 等待信号
	<-c
	fmt.Println("Interrupted, exiting...")
	os.Exit(128 + int(syscall.SIGINT))
}
