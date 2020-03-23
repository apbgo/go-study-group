package main

import (
	"fmt"
	"net/http"
)

// 処理ハンドラ
type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, server.")
}

func main() {
	// ハンドラをエントリポイントと紐付け
	http.Handle("/", Handler{})

	// サーバをlocalhost:8080で起動
	http.ListenAndServe(":8080", nil)
}
