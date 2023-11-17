package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type Rect struct {
	length, width float64
}

func (r Rect) area() float64 {
	return r.length * r.width
}

func (r Rect) perim() float64 {
	return 2*r.length + 2*r.width
}

type Shape interface {
	area() float64
	perim() float64
}

func measure(s Shape) {
	fmt.Println(s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perim())
}

// func measureCircle(s Circle) {
// 	fmt.Println(s)
// 	fmt.Println("Area: ", s.area())
// 	fmt.Println("Perimeter: ", s.perim())
// }

// func measureRect(s Rect) {
// 	fmt.Println(s)
// 	fmt.Println("Area: ", s.area())
// 	fmt.Println("Perimeter: ", s.perim())
// }

func main() {
	fmt.Println("Interface!")
	c := Circle{radius: 1}
	r := Rect{length: 5, width: 2}
	// fmt.Printf("Value = %+v, Type = %T\n", c, c)
	// fmt.Printf("Value = %+v, Type = %T\n", r, r)
	// measureCircle(c)
	// measureRect(r)
	measure(c)
	measure(r)
}
