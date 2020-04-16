package main

import (
	"fmt"
	"net/http"
)

// メインのしたい処理
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello middleware.")
}

func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware1の前処理")
		next.ServeHTTP(w, r)
		fmt.Println("middleware1の後処理")
	}
}

func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware2の前処理")
		next.ServeHTTP(w, r)
		fmt.Println("middleware2の後処理")
	}
}

func middleware3(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware3の前処理")
		next.ServeHTTP(w, r)
		fmt.Println("middleware3の後処理")
	}
}

func middleware(f http.HandlerFunc) http.HandlerFunc {
	return middleware1(middleware2(middleware3(f)))
}

func main() {
	mux := http.NewServeMux()

	// 愚直に書く場合
	//mux.HandleFunc("/", middleware1(middleware2(middleware3(indexHandler))))
	// ちょっとシンプルに(書き方は色々ある)
	mux.HandleFunc("/", middleware(indexHandler))

	http.ListenAndServe(":8080", mux)
}
