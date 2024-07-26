package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Replacement by regex
func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile("b+")
	up := re.ReplaceAllStringFunc(te, strings.ToUpper)

	fmt.Println("up:", up)
}
