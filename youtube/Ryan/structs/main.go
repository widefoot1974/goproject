package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func main() {
	u := User[string]{
		ID:   0,
		Name: "John",
		Data: "Seoul",
	}

	fmt.Printf("user: %+v\n", u)
}
