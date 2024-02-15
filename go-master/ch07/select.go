package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan<- int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
			log.Println("createNumber:", createNumber)
		case <-end:
			log.Println("Ended!")
			return
		case <-time.After(10 * time.Second):
			log.Println("time.After()!")
			return
		}
	}
}

func main() {

	log.SetFlags(log.Lmicroseconds)

	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	log.Printf("Going to create %d random numbers.\n", n)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		log.Printf("%d \n", <-createNumber)
	}

	end <- true
	wg.Wait()
	log.Println("Exiting...")
}
