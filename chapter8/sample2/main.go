package main

import (
	"fmt"
	"time"
)

// nilチャネルが送受信されない例
func main() {
	ch1 := make(chan int)
	var ch2 chan string // nil
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 10
	}()
	go func() {
		// 通らない
		ch2 <- "aaa"
		fmt.Println("ch2にデータが入った")
	}()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}
