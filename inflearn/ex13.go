package main

import (
	"fmt"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func main() {
	m := cmap.New[any]()
	m.Set("a1", 1234)
	m.Set("a2", "morning")
	m.Set("a3", []int{2, 3, 4, 5})
	if v, ok := m.Get("a2"); ok {
		fmt.Printf("a2: %v\n", v)
	}
	m.IterCb(func(key string, value any) {
		fmt.Printf("key: %v, value: %#v\n", key, value)
	})
}
