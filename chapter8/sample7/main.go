package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// WithCancel()で新しいctxとcancel()を返却される
	childCtx, cancel := context.WithCancel(ctx)

	go child(childCtx, "child")

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("parent end")

	time.Sleep(time.Second * 5)
}

func child(ctx context.Context, str string) {
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(str + ": canceled. \n")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println(str)
		}
	}
}
