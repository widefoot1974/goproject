package main

import (
	"fmt"
)

type Updater interface {
	Update()
}

type Transform struct {
	position int
}

type Ememy struct {
	Transform
}

func (e *Ememy) checkTilesCollided() {
	// Logic
	fmt.Println("enermy walking on tile", e.position)
}

func (e *Ememy) Update() {
	e.position += 1
	e.checkTilesCollided()
}

func main() {
	e := &Ememy{}
	for i := 0; i < 10; i++ {
		Update(e)
	}

	fmt.Printf("")
}

func Update(u Updater) {
	u.Update()
}
