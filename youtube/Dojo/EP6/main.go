package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	start := time.Now()
	defer func() {
		log.Printf("elapsed time: %v\n", time.Since(start))
	}()
	log.Println("program is started.")

	var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go taskFunc(i, &wg)
	// }
	// wg.Wait()

	wg.Add(1)
	wg.Done()
	// wg.Done()
	wg.Wait()

	log.Println("program is completed.")
}

func taskFunc(num int, wg *sync.WaitGroup) {

	time.Sleep(time.Second * 2)
	log.Printf("message number is %v\n", num)
	time.Sleep(time.Second * 2)
	wg.Done()
}
