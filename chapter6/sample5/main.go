package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	query = "SELECT * FROM i_user WHERE user_id IN (?, ?)"
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
	if err := Sample5(); err != nil {
		log.Fatal(err)
	}
}

func Sample5() error {
	defer db.Close()

	ctx := context.Background()

	// Transactionを開始
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}
	// 通常はCommit()
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, query, 1, 2)
	if err != nil {
		return err
	}
	// Close()しておかないと次のクエリが発行できないため
	rows.Close()

	// PreparedStmtをトランザクションで利用したい場合
	// tx.Prepare(query)でもPreparedStmtを発行することが可能
	// しかし上記の方法で発行したstmtはtxが終了するとCloseされてしまい再利用出来ない
	// tx.Stmtでシングルトンで定義したstmtを渡すことで、再利用可能
	txStmt := tx.Stmt(stmt)
	defer txStmt.Close()
	rows, err = txStmt.QueryContext(ctx, 1, 2)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
