package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	result := rand.Intn(100)
	fmt.Println("DoWork", result)
	return result
}

func main() {

	startNow := time.Now()
	dataChan := make(chan int)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}

	fmt.Println("This operation took:", time.Since(startNow))
}
