package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ch := make(chan string)
	numRounds := 5

	go throwing(ch, numRounds)

	// for message := range ch {
	// 	fmt.Println(message)
	// }

	for {
		message, open := <-ch
		if !open {
			break
		}
		fmt.Println(message)
	}

	fmt.Println("out of for loop")
	time.Sleep(time.Second * 1)
}

func throwing(ch chan string, numRounds int) {
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		ch <- fmt.Sprintf("Your score is %v", score)
	}
	// close(ch)
}
