package main

import (
	"fmt"
	"log"
	"sync"
)

type singleton struct {
	Value int
}

var (
	once     sync.Once
	instance *singleton
)

func getInstance() *singleton {
	once.Do(func() {
		log.Println("Creating singleton instance...")
		instance = &singleton{Value: 42}
	})
	return instance
}

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			log.Printf("Instance value: %v\n", getInstance().Value)
		}()
	}

	var input string
	fmt.Scanln(&input)
}
