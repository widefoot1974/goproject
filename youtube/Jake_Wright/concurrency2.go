package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Printf("msg = %v\n", msg)

	msg = <-c
	fmt.Printf("msg = %v\n", msg)
}
