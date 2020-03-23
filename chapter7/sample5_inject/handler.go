package main

import (
	"fmt"
	"net/http"
)

type DBHandler struct {
	env string
}

// NewDBHandler コンストラクタ
func NewDBHandler(env string) *DBHandler {
	return &DBHandler{
		env: env,
	}
}

func (h *DBHandler) GetData(w http.ResponseWriter, r *http.Request) {
	// envに設定された環境のDBからデータ取得して返却など
	fmt.Fprint(w, h.env, "環境のデータを取得しました")
}

func (h *DBHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// envに設定された環境のDBのデータを削除など
	fmt.Fprint(w, h.env, "環境のデータを削除しました")
}
