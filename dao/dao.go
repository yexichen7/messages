package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:277187@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&loc=Local&parseTime=True")
	if err != nil {
		log.Fatalf("connect mysql error:%v", err)
		return
	}
	DB = db
	fmt.Println(db.Ping())
}
