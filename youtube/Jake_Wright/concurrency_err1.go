package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	// for {
	// 	msg := <-c
	// 	fmt.Printf("msg = %s\n", msg)
	// }

	for msg := range c {
		fmt.Printf("msg = %s\n", msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("i = %d\n", i)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
