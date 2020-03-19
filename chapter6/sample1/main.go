package main

import (
	"database/sql"
	"log"

	// MySQLを利用するのでDriverをロードする
	"github.com/apbgo/go-study-group/chapter6"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := Sample1(); err != nil {
		log.Fatal(err)
	}
}

func Sample1() error {
	// 接続先設定
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:5446)/chapter6?parseTime=true")
	if err != nil {
		return err
	}
	// アプリケーションが終了するときにCloseするように
	defer db.Close()

	// SQLを実行
	rows, err := db.Query("SELECT * FROM i_user")
	if err != nil {
		return err
	}
	defer rows.Close()

	var records []chapter6.IUser

	// 1レコードずつ、処理する
	for rows.Next() {
		// モデルを作成して、カラムへのポインタをScan()に渡す
		record := chapter6.IUser{}
		if err = rows.Scan(
			&record.UserID,
			&record.OSType,
			&record.Name,
			&record.GamestartDatetime,
			&record.LatestVersion,
			&record.CreatedAt,
			&record.UpdatedAt,
			&record.DeletedAt,
		); err != nil {
			return err
		}
		records = append(records, record)
	}

	// 処理中にエラーが発生する場合もあるのでここの処理を忘れずに
	if err = rows.Err(); err != nil {
		return err
	}

	log.Printf("%+v", records)
	return nil
}
