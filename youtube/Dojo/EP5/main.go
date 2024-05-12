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

	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMsg(ch1, "message1")
	go sendMsg(ch2, "message2")

	select {
	case message := <-ch1:
		log.Println("user1", message)
	case message := <-ch2:
		log.Println("user2", message)
	}

	roughlyFair()

}

func sendMsg(ch chan string, message string) {
	ch <- message
}

func roughlyFair() {
	ch1 := make(chan interface{})
	close(ch1)
	ch2 := make(chan interface{})
	close(ch2)

	var ch1Count, ch2Count int
	for i := 0; i < 1000; i++ {
		select {
		case <-ch1:
			ch1Count++
		case <-ch2:
			ch2Count++
		}
	}

	log.Printf("ch1Count = %v, ch2Count = %v\n", ch1Count, ch2Count)
}
