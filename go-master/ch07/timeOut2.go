package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var result = make(chan bool)

func timeout(t time.Duration) {
	temp := make(chan int)
	go func() {
		log.Println("func() start")
		defer close(temp)
		time.Sleep(3 * time.Second)
	}()

	select {
	case <-temp:
		result <- false
		log.Println("result is false")
	case <-time.After(t):
		result <- true
		log.Println("result is true")
	}
}

func main() {

	log.SetFlags(log.Lmicroseconds)
	log.Println("Started")

	arguments := os.Args
	if len(arguments) != 2 {
		log.Println("Please provide a time duration in milliseconds!")
		return
	}

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	log.Printf("Timeout period is %s\n", duration)

	go timeout(duration)

	val := <-result
	if val {
		log.Println("Time out!")
	} else {
		log.Println("OK")
	}

	log.Println("Ended")
}
