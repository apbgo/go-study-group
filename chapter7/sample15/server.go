package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apbgo/go-study-group/chapter7/sample15/model"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// リクエストBodyの内容を取得
	var req model.Request
	// []byteに変換
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(req)

	// レスポンスの作成
	response := model.Response{
		Status: http.StatusOK,
		Data:   fmt.Sprintf("ID=%v, Name=%s", req.ID, req.Name),
	}
	var res bytes.Buffer
	enc := json.NewEncoder(&res)
	if err = enc.Encode(response); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(res.Bytes())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("localhost:8080", nil)
}
