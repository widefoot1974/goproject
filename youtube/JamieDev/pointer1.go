package main

import (
	"fmt"
)

func main() {

	i, j := 42, 2701

	p := &i
	fmt.Printf("*p = %v\n", *p)
	fmt.Printf("p type = %T\n", p)
	*p = 21
	fmt.Printf("i = %v\n", i)

	p = &j
	*p = *p / 37
	fmt.Printf("j = %v\n", j)
}
