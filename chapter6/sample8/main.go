package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"apb-gitlab.abot.sh/apbgo/golib/pkg/clock"
	"apb-gitlab.abot.sh/apbgo/golib/pkg/db/condition"
	"apb-gitlab.abot.sh/apbgo/golib/pkg/db/orm"
	Orm "apb-gitlab.abot.sh/apbgo/golib/pkg/db/orm"
	"apb-gitlab.abot.sh/apbgo/golib/pkg/logger"
	"github.com/apbgo/go-study-group/chapter6/sample8/db/i_user"
	"github.com/apbgo/go-study-group/chapter6/sample8/db/i_user_item"
)

func main() {
	if err := Sample8_1(); err != nil {
		log.Fatal(err)
	}

	if err := Sample8_2(); err != nil {
		log.Fatal(err)
	}
}

func Sample8_1() error {
	log.Printf("--------------------------- Sample8_1 ---------------------------\n\n\n")
	orm, err := newOrm()
	if err != nil {
		return err
	}
	defer orm.Close()

	ctx := context.Background()
	// インターセプターなどで1リクエスト, 1度実行するイメージ
	// このスコープがトランザクション境界になる
	ctx = orm.WithMultiSession(ctx)
	defer orm.RollbackAll(ctx)

	// PKで1件取得
	iUser := i_user.IUser{
		UserId: 1,
	}
	_, err = orm.Get(ctx, &iUser)
	if err != nil {
		return err
	}

	// 条件で複数件取得
	iUserItems := make(i_user_item.IUserItems, 0)

	// user_id = 1を検索する
	condition, _ := condition.WhereIntEQ("user_id").WithValue(1)
	err = orm.Find(ctx, &iUserItems, condition)
	if err != nil {
		return err
	}
	log.Println(len(iUserItems))

	// Insert
	_, err = orm.Insert(ctx, &i_user_item.IUserItem{
		UserId: 1,
		ItemId: 3,
		Count:  100,
	})
	if err != nil {
		return err
	}

	// Updateは変更されたカラムのみに対して実行される
	iUser.Name = "黒の騎士"
	_, err = orm.Update(ctx, &iUser)
	if err != nil {
		return err
	}

	// Deleteは論理削除対応のテーブルの場合はUpdateが実行される
	// 論理削除されたレコードはSelectされない
	_, err = orm.Delete(ctx, &iUser)
	if err != nil {
		return err
	}
	ok, err := orm.Get(ctx, &iUser)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("論理削除されたレコードがSelectされました")
	}

	// 論理削除対応のテーブルを物理削除する場合
	_, err = orm.Purge(ctx, &iUser)
	if err != nil {
		return err
	}

	return nil
}

func Sample8_2() error {
	log.Printf("--------------------------- Sample8_2 ---------------------------\n\n\n")

	orm, err := newOrm()
	if err != nil {
		return err
	}
	defer orm.Close()

	// キャッシュ可能なORM
	corm := Orm.NewContextCachedOrm(orm)
	ctx := context.Background()
	ctx = corm.WithMultiSession(ctx)
	defer corm.RollbackAll(ctx)

	// 条件で複数件取得
	iUserItems := make(i_user_item.IUserItems, 0)
	condition1, _ := condition.WhereIntEQ("user_id").WithValue(1)
	log.Println("--------------------------- Find1 ---------------------------")
	err = corm.Find(ctx, &iUserItems, condition1)
	if err != nil {
		return err
	}

	// 同じ条件ではSQLは発行されない
	log.Println("--------------------------- Find2 ---------------------------")
	err = corm.Find(ctx, &iUserItems, condition1)
	if err != nil {
		return err
	}

	if actual, expected := len(iUserItems), 2; actual != expected {
		return fmt.Errorf("要素数がおかしい expected %d, actual %d", expected, actual)
	}

	// 今までの検索条件（Condition）から結果が確実に決まる場合はSQLは発行されない
	log.Println("--------------------------- Get1 ---------------------------")
	_, err = corm.Get(ctx, &i_user_item.IUserItem{
		UserId: 1,
		ItemId: 3,
	})
	if err != nil {
		return err
	}
	log.Println("--------------------------- Get2 ---------------------------")
	_, err = corm.Get(ctx, &i_user_item.IUserItem{
		UserId: 1,
		ItemId: 2,
	})
	if err != nil {
		return err
	}

	// Update, Insert, DeleteはSQLが実行されない!
	log.Println("--------------------------- Update1 ---------------------------")
	iUserItem1, iUserItem2 := iUserItems[0], iUserItems[1]
	iUserItem1.Count += 100
	_, err = corm.Update(ctx, iUserItem1)
	if err != nil {
		return err
	}

	log.Println("--------------------------- Insert1 ---------------------------")
	_, err = corm.Insert(ctx, &i_user_item.IUserItem{
		UserId: 1,
		ItemId: 3,
		Count:  100,
	})
	if err != nil {
		return err
	}
	log.Println("--------------------------- Insert2 ---------------------------")
	_, err = corm.Insert(ctx, &i_user_item.IUserItem{
		UserId: 1,
		ItemId: 4,
		Count:  200,
	})
	if err != nil {
		return err
	}
	log.Println("--------------------------- Insert3 ---------------------------")
	_, err = corm.Insert(ctx, &i_user_item.IUserItem{
		UserId: 1,
		ItemId: 5,
		Count:  300,
	})
	if err != nil {
		return err
	}
	log.Println("--------------------------- Delete1 ---------------------------")
	_, err = corm.Delete(ctx, iUserItem2)
	if err != nil {
		return err
	}

	// SQLは実行されないがSelectの結果に反映される
	log.Println("--------------------------- Find3 ---------------------------")
	err = corm.Find(ctx, &iUserItems, condition1)
	if err != nil {
		return err
	}
	if actual, expected := len(iUserItems), 4; actual != expected {
		return fmt.Errorf("要素数がおかしい expected %d, actual %d", expected, actual)
	}

	// 今までの検索条件（Condition）から結果が確実に決まらない場合はSQLが発行される
	// しかし、これまでの更新も考慮された結果が返却される
	log.Println("--------------------------- Find4 ---------------------------")
	condition2, _ := condition.WhereIntsIN("item_id").WithValue([]int64{1, 2, 3, 4, 5})
	count, err := corm.Count(ctx, &i_user_item.IUserItem{}, condition2)
	if err != nil {
		return err
	}
	if actual, expected := count, 6; actual != expected {
		return fmt.Errorf("要素数がおかしい expected %d, actual %d", expected, actual)
	}

	// Commit時にBegin, Insert, Update, Deleteが実行されたあとにCommitされる
	// Insertは自動でBulkInsertに変更される
	log.Println("--------------------------- CommitAll ---------------------------")
	err = corm.CommitAll(ctx)
	if err != nil {
		return err
	}

	// Commit後はキャッシュがクリアされるためまたSQLが発行される
	log.Println("--------------------------- Find5 ---------------------------")
	err = corm.Find(ctx, &iUserItems, condition1)
	if err != nil {
		return err
	}

	return nil
}

func newOrm() (*orm.Orm, error) {
	return Orm.NewOrm(Orm.ConnectionConfig{
		// 水平、垂直分散可能
		"default": {
			(&Orm.DataSource{
				Driver:                "mysql",
				Host:                  "127.0.0.1",
				Port:                  "5446",
				User:                  "root",
				DB:                    "chapter6",
				MaxIdleConnections:    10,
				MaxOpenConnections:    10,
				MaxConnectionLifetime: time.Minute,
				Options: map[string]string{
					"parseTime": "true",
				},
			}).MustOpen(),
		},
	}, func(orm *Orm.Orm) {
		orm.Logger = logger.NewSystemLogger()
		orm.Clock = clock.NewClock()
	})
}
