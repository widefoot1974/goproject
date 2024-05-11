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
	numRounds := 10

	go throwing_(ch)

	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		ch <- fmt.Sprintf("Your score is %v", score)
	}
	log.Println("out of loop")
	time.Sleep(time.Second * 10)
}

func throwing_(ch chan string) {
	for message := range ch {
		time.Sleep(time.Millisecond * 500)
		log.Println(message)
	}
}
