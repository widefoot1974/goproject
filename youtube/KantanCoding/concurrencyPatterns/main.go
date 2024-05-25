package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{}, 2)
	defer func() {
		log.Println("defer: done close")
		close(done)
	}()

	cows := make(chan interface{}, 100)
	// pigs := make(chan interface{}, 100)

	go func() {
		var count = 0
		for {
			count++
			msg := fmt.Sprintf("moo%v", count)
			time.Sleep(time.Millisecond * 1000)
			select {
			case <-done:
				log.Println("\n[func] done channel received, return")
				return
			case cows <- msg:
				log.Printf("[func] cows channle send %v", msg)
			}
		}
	}()

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case cows <- "oink":
	// 		}
	// 	}
	// }()

	wg.Add(1)
	go consumeCows(done, cows)
	// wg.Add(1)
	// go consumePigs(done, pigs)

	// Wait Signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT)
	<-signalCh
	done <- "sigint"
	done <- "sigint"
	wg.Wait()
	log.Println("\nReceived SIGINT. Exiting...\n")
}

// func consumePigs(done <-chan interface{}, pigs <-chan interface{}) {
// 	defer wg.Done()
// 	for pig := range orDone(done, pigs) {
// 		// do some complex logic
// 		fmt.Printf("cow = [%v]\n", pig)
// 		fmt.Printf("cow = [%v]\n", pig)
// 		fmt.Printf("cow = [%v]\n", pig)
// 		fmt.Printf("cow = [%v]\n", pig)
// 		fmt.Printf("cow = [%v]\n", pig)
// 	}
// }

func consumeCows(done <-chan interface{}, cows <-chan interface{}) {
	defer func() {
		log.Println("defer: wg.Done()")
		wg.Done()
	}()
	for cow := range orDone(done, cows) {
		// do some complex logic
		log.Printf("[consumeCows] cow = [%v]\n", cow)
		// log.Printf("cow = [%v]\n", cow)
		// log.Printf("cow = [%v]\n", cow)
		// log.Printf("cow = [%v]\n", cow)
		// log.Printf("cow = [%v]\n", cow)
	}
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer func() {
			log.Println("defer: relayStream close")
			close(relayStream)
		}()
		for {
			select {
			case <-done:
				log.Println("\n[orDone] done channel received, return")
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				log.Printf("[orDone] v = %v\n", v)
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return relayStream
}
