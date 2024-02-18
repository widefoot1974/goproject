package main

import (
	"fmt"
)

func like(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("Type string, value = %#v\n", i.(string))
	case int:
		fmt.Printf("Type int, value = %#v\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
func main() {
	like("Hello")
	like(51)
	like(15.4)
}
