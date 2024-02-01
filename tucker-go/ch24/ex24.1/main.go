package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c (%v) \n", v, v)
	}
}

func PrintNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d (%v)\n", i, i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumber()

	// time.Sleep(3 * time.Second)
}
