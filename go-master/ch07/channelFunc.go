package main

import (
	"fmt"
	"time"
)

func printer(ch chan<- bool) {
	ch <- true
}

func writeToChanel(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
	return
}

func main() {
	c := make(chan int)
	go writeToChanel(c, 10)
	time.Sleep(time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(time.Second)
	close(c)

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	c1 <- 5
	f2(c1, c2)

	fmt.Println("Read (main):", <-c2)
}
