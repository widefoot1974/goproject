package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	start := time.Now()
	defer func() {
		log.Println(time.Since(start))
	}()

	ch := make(chan string, 2)
	ch <- "First message"
	go func(ch chan string) {
		time.Sleep(time.Second * 1)
		ch <- "Third message"
	}(ch)
	ch <- "Second message"

	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	time.Sleep(time.Second * 2)
}
