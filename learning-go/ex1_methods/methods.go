package main

import "fmt"

type Person struct {
	name string
	age  int
}

// func details(p Person) {
// 	fmt.Println(p.name, p.age)
// }

func (p *Person) details(i int) int {
	p.name = "Peter"
	p.age = i
	fmt.Println(p.name, p.age)
	return 0
}

func main() {
	fmt.Println("Method")
	p := Person{
		name: "James",
		age:  25,
	}
	fmt.Printf("p = %+v\n", p)
	// details(p)
	p.details(10)
	fmt.Printf("p = %+v\n", p)
}
