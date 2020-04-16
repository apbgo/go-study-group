package sample6_json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// リクエストパラメタの取得
	id := r.FormValue("id")
	if id == "" {
		fmt.Fprint(w, "パラメータidがない")
		return
	}
	fmt.Fprint(w, "id:", id)

	// リクエストボディの取得
	defer r.Body.Close()
	var req Request
	// io.Readerを実装している
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(req)

	// リクエストヘッダの取得
	contentType := r.Header.Get("Content-Type")
	fmt.Fprintln(w, contentType)
}

func main() {
	// ハンドラーとパスをDefaultServeMuxに登録する
	http.HandleFunc("/", handler)

	// nilの時はDefaultServeMuxが使われる
	http.ListenAndServe("localhost:8080", nil)
}
