package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{}, 10)
	defer func() {
		fmt.Printf("done channel close\n")
		close(done)
	}()
	cows := make(chan interface{}, 10)

	go func() {
		count := 0
		for {
			count++
			time.Sleep(time.Millisecond * 500)
			select {
			case <-done:
				fmt.Printf("\n[go func] done channel received, return\n")
				return
			case cows <- "moo":
				fmt.Printf("[count:%v] cow channel send!!\n", count)
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows, &wg)

	// Wait Signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT)
	<-signalCh
	done <- "sigint"
	done <- "sigint"
	wg.Wait()
	fmt.Println("\nReceived SIGINT. Exiting...\n")
}

func consumeCows(done <-chan interface{}, cows <-chan interface{}, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("wg.Done()")
		wg.Done()
	}()
	for {
		select {
		case <-done:
			fmt.Printf("\n[consumeCows] done channel received, return\n")
			return
		case v, ok := <-cows:
			fmt.Printf("v = %v, ok = %v\n", v, ok)
			if !ok {
				fmt.Printf("cows channel not ok received, return\n")
				return
			}
		}
	}
}
