package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {
			return worker(i)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func worker(i int) error {
	if i == 5 {
		return fmt.Errorf("エラーが発生")
	}
	fmt.Printf("working ... %v\n", i)
	return nil
}
