package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, server.")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// OSからのシグナルを待つ
	go func() {
		// SIGTERM: コンテナが終了する時に送信されるシグナル
		// SIGINT: Ctrl+c
		sigCh := make(chan os.Signal, 1)
		// 受け取るシグナルを指定
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		// チャネルでの待受、シグナルを受け取るまで以降は処理されない
		<-sigCh

		log.Println("start graceful shutdown server.")
		// タイムアウトのコンテキストを設定（後述）
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
			// 接続されたままのコネクションも明示的に切る
			srv.Close()
		}
		log.Println("HTTPServer shutdown.")
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
