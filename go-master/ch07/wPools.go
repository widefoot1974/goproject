package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Result struct {
	job   Client
	squre int
}

var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		squre := c.integer * c.integer
		output := Result{c, squre}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Printf("size = %v\n", size)

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("strconv.Atoi():", err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("strconv.Atoi():", err)
		return
	}

	log.Println("Started")

	go create(nJobs)

	finished := make(chan interface{})
	go func() {
		for d := range data {
			log.Printf("Client ID: id:%d, integer: %d, square: %d\n",
				d.job.id, d.job.integer, d.squre)
		}
		finished <- true
	}()

	var wg sync.WaitGroup
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)

	log.Printf("Finished: %v\n", <-finished)
}
