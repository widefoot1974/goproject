package main

import (
)

func main() {
	initLogger()
	defer logger.Sync()

	Infof("User %s has logged in from IP %s at %d", "Alice", "192.168.1.1", 20230101)
	Warnf("Payment of $%.2f has been processed for account %d", 123.45, 987654)
}

func task() {
}
