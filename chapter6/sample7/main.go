package main

import (
	"database/sql"
	"log"

	"github.com/apbgo/go-study-group/chapter6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	if err := Sample7(); err != nil {
		log.Fatal(err)
	}
}

func Sample7() error {
	engine, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:5446)/chapter6?parseTime=true")
	if err != nil {
		return err
	}
	defer engine.Close()
	// SQLのログを出力する
	engine.ShowSQL(true)

	sess := engine.NewSession()
	// トランザクションを開始
	if err = sess.Begin(); err != nil {
		return err
	}
	defer func() {
		sess.Rollback()
		sess.Close()
	}()

	var users []chapter6.IUser
	// SELECTしてモデルにバインドする
	if err = sess.Where("user_id IN (?, ?)", 1, 2).Find(&users); err != nil {
		return err
	}
	log.Println(users)

	user := users[0]
	user.Name = sql.NullString{
		String: "キリトLv1000",
		Valid:  true,
	}

	// 更新する
	count, err := sess.Where("user_id = ?", user.UserID).Update(user)
	if err != nil {
		return err
	}
	log.Println(count)

	var joinedUsers []chapter6.JoinedUser
	err = sess.Table("i_user").
		Select("i_user.*, i_user_item.*").
		Join("INNER", "i_user_item", "i_user.user_id = i_user_item.user_id").
		Where("i_user.user_id = ?", 1).
		Find(&joinedUsers)
	if err != nil {
		return err
	}
	log.Printf("%+v", joinedUsers)

	return nil
}
