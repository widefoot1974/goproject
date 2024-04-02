package main

import (
	"fmt"
)

func sum(data ...float64) float64 {
	fmt.Println(data)
	for i, v := range data {
		fmt.Printf("[%v] %v\n", i, v)
	}

	return 0.0
}

func main() {

	a := make([]float64, 5)
	b := []float64{}

	fmt.Println(a)
	fmt.Println(b)

	sum(1.0, 2.0, 3.0)
	sum(0.0)
}
