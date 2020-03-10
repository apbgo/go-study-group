package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	query = "UPDATE i_user set name = ? WHERE user_id = ?"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:5446)/chapter6?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := Sample6(); err != nil {
		log.Fatal(err)
	}
}

func Sample6() error {
	defer db.Close()

	ctx := context.Background()

	// Transactionを開始
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txStmt := tx.Stmt(stmt)
	defer txStmt.Close()

	// UPDATE, INSERT, DELETEなどはExecを利用する
	result, err := txStmt.ExecContext(ctx, "黒の騎士", 1)
	if err != nil {
		return err
	}

	// 更新したレコード数を受け取ることが可能
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println(count)
	return nil
}
