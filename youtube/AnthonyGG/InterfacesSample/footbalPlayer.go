package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall() int
	Name() string
}

type CR7 struct {
	name    string
	stamina int
	power   int
	SUI     int
}

func (f CR7) KickBall() int {
	return f.stamina + f.power*f.SUI
}

func (f CR7) Name() string { return "CR7" }

type Messi struct {
	name    string
	stamina int
	power   int
	SUI     int
}

func (f Messi) KickBall() int {
	return f.stamina + f.power*f.SUI
}

func (f Messi) Name() string { return "Messi" }

type FootbalPalyer struct {
	stamina int
	power   int
}

func (f FootbalPalyer) KickBall() int {
	return f.stamina + f.power
}

func (f FootbalPalyer) Name() string { return "FootbalPalyer" }

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = FootbalPalyer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-2] = CR7{
		stamina: 5,
		power:   5,
		SUI:     5,
	}
	team[len(team)-1] = Messi{
		stamina: 5,
		power:   5,
		SUI:     5,
	}
	for i := 0; i < len(team); i++ {
		shot := team[i].KickBall()
		fmt.Printf("%v is kicking the ball %d\n", team[i].Name(), shot)
	}
}
