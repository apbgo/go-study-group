package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB

	// PrepareStmtをシングルトンでクエリごとに持つことで、再利用することが可能
	// MySQLはコネクションに対してPreparedStmtが紐づいている
	// そのため一見コネクションごとにStmtを持つ必要があるが、GoはPreparedStmtのライフライクルを管理してくれるため
	// クエリごとに1つ作成すればよい
	stmt *sql.Stmt
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:5446)/chapter6?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// コネクションプールはデフォルトで有効になっているが以下の3つの設定を調整することが可能
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Minute * 30)

	stmt, err = db.Prepare("SELECT * FROM i_user WHERE user_id IN (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := Sample4(); err != nil {
		log.Fatal(err)
	}
}

func Sample4() error {
	defer db.Close()

	ctx := context.Background()
	rows, err := stmt.QueryContext(ctx, 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return nil
}
