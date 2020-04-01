package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	quit := make(chan int)

	go func() {
		ch <- 1
	}()
	go func() {
		time.Sleep(time.Second * 5)
		quit <- 0
	}()

	// ラベルなんかも時々使う（これまで触れてなかったので紹介）
LOOP:
	for {
		select {
		case v1 := <-ch:
			fmt.Println(v1)
		case <-quit:
			break LOOP
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("...")
		}
	}
	fmt.Println("finish.")
}
