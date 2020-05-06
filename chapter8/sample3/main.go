package main

import (
	"fmt"
	"sync"
	"time"
)

var mux sync.RWMutex
var data string

// sync.RWMutexのサンプル
func main() {
	data = "aaa"
	mux = sync.RWMutex{}

	go read()
	go read()
	go write()
	go read()

	time.Sleep(time.Second * 5)
}

func read() {
	fmt.Println("read start")
	mux.RLock()
	defer mux.RUnlock()

	time.Sleep(time.Second)
	fmt.Println("read finish")
}

func write() {
	fmt.Println("write start")
	mux.Lock()
	defer mux.Unlock()

	time.Sleep(time.Second * 2)
	data += "bbb"
	fmt.Println("write finish")
}
