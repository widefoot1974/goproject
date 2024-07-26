package main

import (
	"fmt"
	"regexp"
)

// Location by regex
func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile("b+")
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Printf("mm = [%v]\n", mm)
	fmt.Printf("id = [%v]\n", id)

	for _, d := range id {
		fmt.Println(te[d[0]:d[1]])
	}
}
