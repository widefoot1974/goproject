package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("You ars using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go Version", runtime.Version())

	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

}
