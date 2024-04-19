package main

import (
	"fmt"
	"log"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	timeout := time.After(time.Second * 10)
	for {
		select {
		case <-snapshotTicker:
			taskSnapshot()
		case <-saveAfter:
			saveSnapshot()
		case <-timeout:
			log.Printf("Return")
			return
		default:
			// waitForData()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func taskSnapshot() {
	log.Println("Tasking a backup snapshot...")
}

func saveSnapshot() {
	log.Println("All backups saved!")
}

func waitForData() {
	log.Println("Noting to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(time.Second * 1)
	log.Println("time.Tick() After")
	saveAfter := time.After(time.Second * 5)
	log.Println("time.After() After")
	saveBackups(snapshotTicker, saveAfter)
	log.Println("saveBackups() After")
	fmt.Println("==============================")
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	test()
}
