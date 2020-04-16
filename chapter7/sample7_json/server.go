package sample7_json

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンスの作成
	response := Response{
		Status: http.StatusOK,
		Data:   "Hello Server.",
	}
	res, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("token", "tokenなんかもつけられるよ")

	// レスポンスの書き込み
	w.Write(res)
}

func main() {
	// ハンドラーとパスをDefaultServeMuxに登録する
	http.HandleFunc("/", JsonHandler)

	// nilの時はDefaultServeMuxが使われる
	http.ListenAndServe("localhost:8080", nil)
}
