package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipienName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf(
		"Hi %s, it is your birthday on %s",
		bm.recipienName,
		bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(
		`Your "%s" report is ready. You've send %v messages.`,
		sr.reportName,
		sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("============================")
}

func main() {

	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipienName: "John Doe",
		birthdayTime: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
}
