package main

import (
	"fmt"
)

func doMath(f func(int) int, nums []int) []int {
	var results []int

	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	allNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)
}
