package main

import (
	"fmt"
)

type Attacker interface {
	Attack() error
}

type StrongAttacker struct{}
type SuperStrongAttacker struct{}

func (sa StrongAttacker) Attack() error {
	fmt.Println("Wooom, what a strong attack!!")
	return nil
}

func (sa SuperStrongAttacker) Attack() error {
	fmt.Println("OOF, what a super strong attack!!")
	return nil
}

func DoAttack(a Attacker) error {
	return a.Attack()
}

func main() {

	DoAttack(StrongAttacker{})
	DoAttack(SuperStrongAttacker{})
}
