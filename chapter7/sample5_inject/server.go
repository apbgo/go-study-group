package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	masterHandler *DBHandler
}

func InjectDevelop() *App {
	return &App{
		masterHandler: NewDBHandler("dev"),
	}
}

func InjectLocal() *App {
	return &App{
		masterHandler: NewDBHandler("local"),
	}
}

func main() {
	// 環境変数でenvを受け取る
	env := os.Getenv("APP_ENV")

	// DI 設定を注入（今回はenvを切り替えている）
	// DIできるようにAppというstructにHandlerやServiceを格納するテクニック
	// 下のIndexHandlerなどは外から値や実体を与えて処理を変えるような事はできない
	var app *App
	switch env {
	case "dev":
		app = InjectDevelop()
	default:
		app = InjectLocal()
	}

	// API群
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/get", app.masterHandler.GetData)
	http.HandleFunc("/delete", app.masterHandler.Delete)

	// サーバ起動
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Server.")
}
