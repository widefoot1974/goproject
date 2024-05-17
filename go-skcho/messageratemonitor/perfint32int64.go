package main

import (
	"fmt"
	"time"
)

func testInt32() {
	start := time.Now()
	var sum int32 = 0
	for i := int32(0); i < 100000000; i++ {
		sum += i
	}
	fmt.Printf("Int32 sum: %d, took: %s\n", sum, time.Since(start))
}

func testInt64() {
	start := time.Now()
	var sum int64 = 0
	for i := int64(0); i < 100000000; i++ {
		sum += i
	}
	fmt.Printf("Int64 sum: %d, took: %s\n", sum, time.Since(start))
}

func main() {
	testInt32()
	testInt64()
}
