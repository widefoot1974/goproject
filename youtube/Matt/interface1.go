package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
	fmt.Printf("%T %[1]v\n", s)
}
