package main

import (
	"fmt"
	"sync"
)

func main() {
	var cm sync.Map
	cm.Store("c1", 689)
	cm.Store("c2", "Have a nice day")
	cm.Store(123, []int{1, 2, 3})

	checkAndPrint(&cm, "c3")
	checkAndPrint(&cm, "c1")
	checkAndPrint(&cm, 123)

}

// 반복되는 로직을 함수로 분리하여 가독성 개선
func checkAndPrint(cm *sync.Map, key interface{}) {
	if v, ok := cm.Load(key); ok {
		fmt.Printf("%v: %v\n", key, v)
	} else {
		fmt.Printf("%v does not exist\n", key)
	}
}
