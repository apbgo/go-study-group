package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := sample14(); err != nil {
		log.Fatal(err)
	}
}

func sample14() error {
	// クライアントにタイムアウトを設定
	client := http.Client{
		Timeout: time.Second * 5,
	}
	// タイムアウト設定したコンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		return err
	}
	// リクエストにコンテキストを設定
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
