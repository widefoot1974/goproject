package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Input goroutines counter!")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d goroutines.\n", count)

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
