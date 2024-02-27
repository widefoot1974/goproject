package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Printf("msg = %v\n", msg)
	// }

	for msg := range c {
		fmt.Printf("msg = %v\n", msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
