package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("実行終了")
			return
		default:
			fmt.Println("実行中")
			time.Sleep(time.Second)
		}
	}
}
