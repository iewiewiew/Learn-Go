package utils

import (
	"database/sql"
	"fmt"
	"log"
	"mix.com/mix/config"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * @author       weimenghua
 * @time         2024-07-12 10:13
 * @description
 */

// MySQLClient 封装了 MySQL 的操作
type MySQLClient struct {
	DB *sql.DB
}

// NewMySQLClient 创建一个新的 MySQLClient 实例
func NewMySQLClient(user, password, host, dbname string) (*MySQLClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLClient{DB: db}, nil
}

// Query 执行查询并返回结果
func (client *MySQLClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return client.DB.Query(query, args...)
}

// Exec 执行命令（如插入、更新、删除）
func (client *MySQLClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return client.DB.Exec(query, args...)
}

// Close 关闭数据库连接
func (client *MySQLClient) Close() error {
	return client.DB.Close()
}

func main() {
	// 替换为实际的数据库连接信息
	user := config.Conf.Database.UserName
	password := config.Conf.Database.Password
	host := config.Conf.Database.Host
	dbname := config.Conf.Database.DBName

	client, err := NewMySQLClient(user, password, host, dbname)
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
