package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process Id: %d\n", os.Getpid())
	sigsCh := make(chan os.Signal, 1)

	signal.Notify(sigsCh)

	start := time.Now()
	go func() {
		for {
			sig := <-sigsCh
			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execute time:", duration)
			case syscall.SIGUSR1:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("Caught:", sig)
			}
		}
	}()

	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}
}
