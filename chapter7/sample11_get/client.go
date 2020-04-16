package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	if err := sample11(); err != nil {
		log.Fatal(err)
	}
}

func sample11() error {
	values := url.Values{}
	values.Add("id", "100")

	// 最も単純なGETの例 -----------------------
	res, err := http.Get("http://localhost:8080" + "?" + values.Encode())
	if err != nil {
		return err
	}
	// Responseは必ずClose
	defer res.Body.Close()
	fmt.Println(res)

	// Clientを使って詳細な設定をしたGETの例----------------
	client := http.Client{
		// タイムアウトなど色々設定できる
		//Transport:     nil,
		//CheckRedirect: nil,
		//Jar:           nil,
		//Timeout:       0,
	}
	// Request を生成
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = values.Encode()
	// ヘッダーの付与
	req.Header.Add("Token", "Tokenなんかも付けられるよ")
	// Client.DoでGET投げる
	res, err = client.Do(req)
	if err != nil {
		return err
	}
	// responseは必ずClose
	defer res.Body.Close()
	fmt.Println(res)

	return nil
}
