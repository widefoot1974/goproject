package main

import "fmt"

type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance     float64
	fuel         float64
	averagespeed string
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func totalMileage(m MotorVehicle) {
	au := m.(BMW)
	fmt.Println(au.averagespeed)
}

func main() {
	b1 := BMW{
		distance:     167.9,
		fuel:         36,
		averagespeed: "58",
	}

	totalMileage(b1)
}
