package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := Sample3(); err != nil {
		log.Fatal(err)
	}
}

func Sample3() error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:5446)/chapter6?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()

	// Goの標準のコンテキスト
	// Webアプリケーションの場合はリクエストに対して1つ以上のコンテキストオブジェクトを持ち回すことが多い
	// https://qiita.com/marnie_ms4/items/985d67c4c1b29e11fffc
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Nanosecond*1)
	defer cancel()

	// SQLパッケージでは、コンテキストは実行中のキャンセルのために利用される
	// この例では1ns後にタイムアウトするコンテキストを渡すため、エラーになるはず
	rows, err := db.QueryContext(ctx, "SELECT * FROM i_user WHERE user_id IN (?, ?)", 1, 2)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
