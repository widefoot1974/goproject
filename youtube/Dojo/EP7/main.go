package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var count int

func main() {
	log.SetFlags(log.Lmicroseconds)
	start := time.Now()
	defer func() {
		log.Printf("elapsed time: %v\n", time.Since(start))
	}()
	log.Println("program is started.")

	iterations := 1000000

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}
	wg.Wait()

	fmt.Printf("Resulted count is: %v\n", count)
}

func increment(wg *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	count++
	wg.Done()
}
