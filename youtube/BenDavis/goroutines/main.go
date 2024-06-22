package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			log.Println("IDX from first func:", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			log.Println("IDX from second func:", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	wg.Wait()
	log.Println("Done")
}
