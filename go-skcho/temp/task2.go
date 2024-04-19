package main

import (
	"fmt"
)

func task1(a *[5]int) {
	(*a)[0] = 100
	(*a)[1] = 200
	(*a)[2] = 300
}

func task2(s *[]int) {
	*s = append(*s, 1)
	*s = append(*s, 2)
	*s = append(*s, 3)
}

func main() {

	var array1 [5]int
	var slice1 []int

	fmt.Printf("array1 = %#v\n", array1)
	fmt.Printf("slice1 = %#v\n", slice1)

	task1(&array1)

	task2(&slice1)

	fmt.Printf("array1 = %#v\n", array1)
	fmt.Printf("slice1 = %#v\n", slice1)

	fmt.Printf("")
}
