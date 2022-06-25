package utils

/**
 * @author       weimenghua
 * @time         2023-08-15 13:48
 * @description  Go MySQL 增删改查
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
)

func go_mysql() {
	// 连接 MySQL 数据库
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功!")

	// 执行查询
	rows, err := db.Query("select id, name from dbname.t_table_info")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 插入数据 @todo 没有生成随机数
	id := rand.Intn(100)
	fmt.Println("随机数: ", id)

	insertStmt, err := db.Prepare("insert into dbname.t_table_info(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(id, "Go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据插入成功!")
}

func main() {
	go_mysql()
}
