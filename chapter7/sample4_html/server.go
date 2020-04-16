package main

import (
	"io"
	"log"
	"net/http"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>HTMLサンプルページ</title>
</head>
<body>
<h1>HTMLサンプルです</h1>
当然ベタなHTMLだけではなくて、html/templateなどのライブラリもあるよ</br>
<a href="https://blog.y-yuki.net/entry/2017/07/04/100000" target="_blank">参考：テンプレートの使い方</a>
</body>
</html>
`)
}

func main() {
	http.HandleFunc("/", htmlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
