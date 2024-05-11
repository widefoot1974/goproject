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

	Users := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, user := range Users {
		go attack(user)
	}

	time.Sleep(time.Second * 2)

}

func attack(target string) {
	fmt.Println("Throwing user stars at", target)
	time.Sleep(time.Second * 1)
}
