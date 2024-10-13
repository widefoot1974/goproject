package main

import (
	"goproject/inflearn/cmap"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	c := cmap.New()
	c.Set("a1", 32)
	c.Set("a2", "Good Evening")
	spew.Dump(c.Get("a2"))
}
