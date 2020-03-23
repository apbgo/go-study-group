package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	if err := sample13(); err != nil {
		log.Fatal(err)
	}
}

func sample13() error {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		return err
	}
	// Responseは必ずClose
	defer res.Body.Close()
	var response Response
	// res.Bodyはio.Readerを実装している
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
