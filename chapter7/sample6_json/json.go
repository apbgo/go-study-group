package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	// jsonタグでJSONのフィールド名を指定する
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{
		Name: "sampleName",
		Age:  10,
	}

	// エンコード1
	buf1, err := json.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf1))

	// エンコード2
	var buf2 bytes.Buffer
	enc := json.NewEncoder(&buf2)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf2.String())

	// デコード1
	var p2 Person
	if err := json.Unmarshal(buf1, &p2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)

	// デコード2
	var p3 Person
	dec := json.NewDecoder(&buf2)
	if err := dec.Decode(&p3); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p3)
}
