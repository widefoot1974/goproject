package main

import "log"

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func main() {
	a := 4
	squreVal(a)
	log.Printf("a = %v, &a = %v\n", a, &a)

	sqaureAdd(&a)
	log.Printf("a = %v, &a = %v\n", a, &a)
}

func squreVal(v int) {
	log.Printf("v = %v\n", v)
	v *= v
	log.Printf("v = %v, &v = %v\n", v, &v)
}

func sqaureAdd(p *int) {
	*p *= *p
	log.Printf("p = %v, *p = %v\n", p, *p)
}
