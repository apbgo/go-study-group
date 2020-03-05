package chapter5

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
func cut() {
	flag.Parse()

	// このValidationを関数1つ目に切り出す ---------
	// ヒント：flagの内容を渡してやって、バリデーションし、エラーがあれば返すような関数にできる
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "ファイルパスを指定してください。")
		os.Exit(1)
	}
	if *fields < 0 {
		fmt.Fprintln(os.Stderr, "-f は1以上である必要があります")
		os.Exit(1)
	}
	// ---------------------------------------

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// この部分をCutコマンドとして関数2つ目に切り出す------
	// ヒント：NewScannerにfileを渡しているが、NewScannerはio.Readerであれば何でも良い
	// また、出力も現在fmt.Println(s)にしているが、io.Writerを使って書き出す先を指定できるようにしてやる
	// 関数の引数で読み出すio.Readerと、
	// 書き出すio.Writer (本関数からはos.Stdout, テストからはbyte.Bufferなどへ)を指定できるようにすると良い
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		sb := strings.Split(text, *delimiter)
		if len(sb) < *fields-1 {
			fmt.Fprintln(os.Stderr, "-fの値に該当するデータがありません")
			os.Exit(1)
		}
		s := sb[*fields-1]
		fmt.Println(s)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// ------------------------------------------------
}
