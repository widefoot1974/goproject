package main

import (
	"log"
	"time"
)

func main() {

	log.SetFlags(log.Lmicroseconds)
	log.Println("Started")

	c1 := make(chan string)
	go func() {
		log.Println("func1() start")
		time.Sleep(3 * time.Second)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		log.Printf("res1 = %v\n", res)
	case <-time.After(time.Second):
		log.Println("timeout c1")
	}

	c2 := make(chan string)
	go func() {
		log.Println("func2() start")
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		log.Printf("res2 = %v\n", res)
	case <-time.After(4 * time.Second):
		log.Println("timeout c2")
	}

	log.Println("Ended")
}
