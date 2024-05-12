package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	start := time.Now()
	defer func() {
		log.Println(time.Since(start))
	}()

	ch := make(chan string)
	numRounds := 5

	go throwing(ch, numRounds)

	for {
		message, open := <-ch
		if !open {
			log.Println("break")
			break
		}
		log.Println(message)
	}

	log.Println("out of for loop")
	time.Sleep(time.Second * 1)
}

func throwing(ch chan string, numRounds int) {
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		ch <- fmt.Sprintf("Your score is %v", score)
	}
	close(ch)
}
