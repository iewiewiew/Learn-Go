package main

import (
	"fmt"
	"log"
	"mix.com/mix/config"
	"mix.com/mix/utils"
)

/**
 * @author       weimenghua
 * @time         2024-07-12 14:27
 * @description
 */

func main() {
	// 替换为实际的数据库连接信息
	user := config.Conf.Database.UserName
	log.Println(user)
	password := config.Conf.Database.Password
	host := config.Conf.Database.Host
	dbname := config.Conf.Database.DBName

	client, err := utils.NewMySQLClient(user, password, host, dbname)
	if err != nil {
		log.Fatalf("Failed to create MySQL client: %v", err)
	}
	defer client.Close()

	// 示例查询
	rows, err := client.Query("SELECT * FROM user")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var user_name string
		var user_age string
		if err := rows.Scan(&id, &user_name, &user_age); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("User ID: %d, UserName: %s, UserAge: %s\n", id, user_name, user_age)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Rows error: %v", err)
	}
}
