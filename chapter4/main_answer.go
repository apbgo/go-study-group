package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var delimiter = flag.String("d", ",", "区切り文字を指定してください")
var fields = flag.Int("f", 1, "フィールドの何番目を取り出すか指定してください")

// go-cutコマンドを実装しよう
func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "ファイルパスを指定してください。")
		os.Exit(1)
	}

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// bufio.NewWriterを使えばもっと効率の良い書き方ができるかも？
	// 興味のある人は調べてみましょう！
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, *delimiter)[*fields-1]
		fmt.Println(s)
	}
}
