package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if !e.isSubscriberd {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscriberd bool
	body          string
}

func print(p printer) {
	p.print()
}

func task(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("========================================")
}

func main() {
	e := email{
		isSubscriberd: false,
		body:          "I want my monony back",
	}
	task(e, e)
	fmt.Printf("")
}
