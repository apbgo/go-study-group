package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := Sample2(); err != nil {
		log.Fatal(err)
	}
}

func Sample2() error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:5446)/chapter6?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()

	// パラメータを渡す場合は、第2引数以降
	// Query()はPreparedStatementを発行、SQLを実行、PreoaredStatementのクローズを行う
	// SQLインジェクションの心配はない
	rows, err := db.Query("SELECT * FROM i_user WHERE user_id IN (?, ?)", 1, 2)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
