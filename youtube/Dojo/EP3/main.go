package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ch := make(chan string, 2)
	ch <- "First message"
	go func(ch chan string) {
		time.Sleep(time.Second * 1)
		ch <- "Third message"
	}(ch)
	ch <- "Second message"

	fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	time.Sleep(time.Second * 2)
}
