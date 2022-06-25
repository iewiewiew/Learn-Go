package main

/**
 * @author       weimenghua
 * @time         2024-07-15 18:44
 * @description
 */

/**
在 Go 语言中，import 语句用于导入其他包中的代码，称为匿名导入（blank import）或空白导入。。有时候可能只需要导入包中的一部分代码，而不是全部代码。在这种情况下，可以使用下划线 _ 替代导入路径中的标识符，以表示忽略该标识符。
*/
import (
	_ "fmt"
	_ "github.com/go-sql-driver/mysql" // 使用 _ 来表示导入的包中的标识符不需要被使用
)

func main() {
	// fmt.Println("Hello, world!") // 编译错误：undefined：fmt.Println
}
