package main

import (
	"fmt"
)

// Define a struct that holds the accumulated lines and provides methods for adding lines and retrieving them.
type Concatter struct {
	lines string
}

// NewConcatter creates and returns a new Concatter instance.
func NewConcatter() *Concatter {
	return &Concatter{lines: ""}
}

// AddLine adds a new line to the Concatter's lines based on the provided name and value.
func (c *Concatter) AddLine(name string, value interface{}) {
	c.lines += fmt.Sprintf(" %-10v : %v\n", name, value)
}

func (c *Concatter) SrtLine(name string, value interface{}) {
	c.lines += fmt.Sprintf("==========[ %v ]==========\n", name)
}

func (c *Concatter) EndLine(name string, value interface{}) {
	c.lines += fmt.Sprintf(" %-10v : %v\n", name, value)
	c.lines += fmt.Sprintln("==============================")
}

// String returns the accumulated lines.
func (c *Concatter) String() string {
	return c.lines
}

func main() {
	log := NewConcatter()
	log.SrtLine("Header", nil)
	log.AddLine("SrcSubject", "EIF-G-0")
	log.AddLine("SrcId", 1)
	log.EndLine("SrcId", 1)

	// Now you can directly print the accumulated lines without adding a dummy line.
	fmt.Println(log.String())
}
