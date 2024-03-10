package main

import (
	"fmt"
	"time"
)

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {

			ch <- struct{}{}
			fmt.Printf("Databases %v is online\n", i+1)

		}
	}()
	return ch
}

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		time.Sleep(time.Second * 1)
		<-dbChan
	}
}

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10)
	fmt.Println("All databases are online!")

}

func main() {

	test(5)
}
