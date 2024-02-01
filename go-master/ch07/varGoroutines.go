package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Input goroutines counter!")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup
	fmt.Printf("waitGroup = %#v\n", waitGroup)

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("waitGroup = %#v\n", waitGroup)
	waitGroup.Wait()

	fmt.Println("\nExiting...")
}
