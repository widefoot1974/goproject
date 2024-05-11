package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		log.Println(time.Since(start))
	}()

	signalCh := make(chan bool)
	user := "Tommy"
	go attack(user, signalCh)

	attackResult := <-signalCh
	log.Printf("attackResult = %v\n", attackResult)
}

func attack(target string, signalCh chan bool) {
	log.Println("Throwing user stars at", target)
	time.Sleep(time.Second * 1)
	signalCh <- true
}
