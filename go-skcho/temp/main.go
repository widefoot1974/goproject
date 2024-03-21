package main

import (
	"fmt"
)

func main() {

	fmt.Printf("")

	// ages := make(map[string]int)
	// ages["John"] = 37
	// ages["Mary"] = 24
	// ages["Mary"] = 21 //overwrites 24

	ages := map[string]int{
		"widefoot": 51,
	}

	fmt.Printf("ages = %v\n", ages)
	fmt.Printf("ages = %#v\n", ages)
	fmt.Printf("ages = %T\n", ages)

}