package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var missingCompleted bool

func main() {

	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			if foundTreasure() {
				once.Do(markMissingCompleted)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	checkMissingCompletion()
}

func checkMissingCompletion() {
	if missingCompleted {
		fmt.Println("Missing is now completed.")
	} else {
		fmt.Println("Missing was a failure.")
	}
}

func markMissingCompleted() {
	missingCompleted = true
	fmt.Println("Marking mission completed.")
}

func foundTreasure() bool {
	num := rand.Intn(10)
	fmt.Printf("num = %v\n", num)
	if num == 0 {
		return true
	} else {
		return false
	}
}
