package main

import (
	"fmt"
	"mix.com/mix/config"
)

/**
 *@author       weimenghua
 *@time         2023-02-25 17:17
 *@description  go example
 */

func main() {
	//fmt.Println("Hello World!")

	//fmt.Println(Green("Color123"))
	//fmt.Println(Blue("Color123"))
	//fmt.Println(Yellow("Color123"))
	//fmt.Println(Cyan("Blue123"))
	//fmt.Println(Red("Blue123"))
	//fmt.Println(Magenta("Blue123"))

	fmt.Println("Database User:", config.Conf.Database.UserName)

	//fmt.Println(utils.GetCurrentDir())
}
