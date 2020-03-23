package main

import (
	"fmt"
	"net/http"
)

// 処理ハンドラ
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, server.")
}

func main() {
	// ServeMuxを新しく作成
	mux := http.NewServeMux()
	// ハンドラをエントリポイントと紐付け
	mux.HandleFunc("/", handler)
	// サーバを作成し、muxを登録（DefaultServeMuxは使わない）
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
	}
	// サーバを起動
	srv.ListenAndServe()
}
