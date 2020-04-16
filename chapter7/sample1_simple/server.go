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
	// ハンドラをエントリポイントと紐付け
	http.HandleFunc("/", handler)

	// サーバをlocalhost:8080で起動
	http.ListenAndServe(":8080", nil)
}
