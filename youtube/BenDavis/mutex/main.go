package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int
}

func (s *SafeCounter) Add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NumMap["key"] = num
	s.NumMap["sum"] += num

}

func main() {
	s := SafeCounter{NumMap: make(map[string]int)}
	var wg sync.WaitGroup

	s.NumMap["sum"] = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("key", s.NumMap["key"])
	fmt.Println("sum", s.NumMap["sum"])
}
