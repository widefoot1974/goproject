package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	user := "Tommy"
	go attack(user)

	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Println("Throwing user stars at", target)
	time.Sleep(time.Second * 1)
}
