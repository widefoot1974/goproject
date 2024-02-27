package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			log.Println(msg1)
		case msg2 := <-c2:
			log.Println(msg2)
		}
	}

}
