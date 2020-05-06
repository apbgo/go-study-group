package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var num int64
	for i := 0; i < 50; i++ {
		go func() {
			// これは競合する可能性がある
			// num++

			atomic.AddInt64(&num, 1)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(num)
}
