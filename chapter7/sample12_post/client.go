package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	if err := sample12(); err != nil {
		log.Fatal(err)
	}
}

func sample12() error {
	values := url.Values{}
	values.Add("id", "100")

	// 最も単純なPOSTの例 -------------------
	res, err := http.PostForm("http://localhost:8080/", values)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Println(res)

	// Clientを使って詳細な設定をしたGETの例----------------
	client := http.Client{
		//Transport:     nil,
		//CheckRedirect: nil,
		//Jar:           nil,
		//Timeout:       0,
	}
	// Request を生成
	req, err := http.NewRequest("POST",
		"http://localhost:8080/", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Client.DoでPOST投げる
	res, err = client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Println(res)

	return nil
}
