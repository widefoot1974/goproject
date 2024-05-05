package main

import "log"

type person struct {
	name string
	age  int
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func initPerson() *person {
	m := person{name: "noname", age: 50}
	log.Printf("initPerson --> %p\n", &m)
	log.Printf("m = %v, &m = %p\n", m, &m)
	return &m
}

func main() {
	m := initPerson()
	log.Printf("m = %v, m = %p\n", m, m)
	log.Printf("main --> %p\n", m)

	m.name = "John"
	log.Printf("m = %v, m = %p\n", m, m)
}
