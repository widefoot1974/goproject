package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		now := time.Now()
		fmt.Printf("Minute = %v, Second = %v\n", now.Minute(), now.Second())
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("")
}
