package main

import (
	"fmt"
)

func main() {

	fmt.Printf("")

	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}

	fmt.Println(myCar)
	fmt.Printf("%p\n", &myCar)
	fmt.Printf("%p\n", myCar)

}
