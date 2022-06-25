package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"log"
)

/**
 * @author       weimenghua
 * @time         2024-07-20 09:45
 * @description
 */

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := goose.Up(db, "./migrations"); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
