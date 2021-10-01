package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SmsCaptcha struct {
	ID        int64  `db:"id"`
	Type      string `db:"type"`
	Phone     string `db:"phone"`
	Code      string `db:"code"`
	Status    int    `db:"status"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type User struct {
	ID        int64  `db:"id"`
	Username  string `db:"username"`
	Phone     string `db:"phone"`
	Avatar    string `db:"avatar"`
	Token     string `db:"token"`
	Status    int    `db:"status"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// // 定义一个全局对象db
// var db *sql.DB

func getDB() (*sql.DB, error) {

	config, _ := getConfig("./config.yaml")

	dsn := config.DB.Username + ":" + config.DB.Password + "@tcp(" + config.DB.Host + ":" + string(config.DB.Port) + ")/" + config.DB.DBName + "?charset=utf8mb4&parseTime=True"
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

type QueryCondition struct {
	Where  string
	Params []string
}

//TODO
func queryOne() {

}
