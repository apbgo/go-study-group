package chapter5

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type cutEntity struct {
	args      []string
	delimiter string
	fields    int
}

//var delimiter = flag.String("d", ",", "区切り文字を指定してください")
//var fields = flag.Int("f", 1, "フィールドの何番目を取り出すか指定してください")

// go-cutコマンドを実装しよう
func main() {
	flag.Parse()

	cutEntity := cutEntity{
		args:      flag.Args(),
		delimiter: *delimiter,
		fields:    *fields,
	}
	err := Validate(cutEntity)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open(cutEntity.args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = Cut(cutEntity, file, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func Validate(entity cutEntity) error {
	if len(entity.args) == 0 {
		return fmt.Errorf("ファイルパスを指定してください。")
	}
	if entity.fields < 0 {
		return fmt.Errorf("-f は1以上である必要があります")
	}
	return nil
}

func Cut(entity cutEntity, r io.Reader, w io.Writer) error {
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	scanner := bufio.NewScanner(r)
	d := []byte(entity.delimiter)
	for scanner.Scan() {
		sb := bytes.Split(scanner.Bytes(), d)
		if len(sb) < entity.fields {
			return fmt.Errorf("-fの値に該当するデータがありません。")
		}
		b := sb[entity.fields-1]
		_, err := bw.Write(b)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
