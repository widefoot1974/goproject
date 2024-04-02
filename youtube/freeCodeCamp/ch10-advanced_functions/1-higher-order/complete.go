package main

import (
	"fmt"
)

func getFormattedMessage(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

func test(messages []string, formatter func(string) string) {
	defer fmt.Println("====================================")
	formattedMessage := getFormattedMessage(messages, formatter)
	if len(formattedMessage) != len(messages) {
		fmt.Println("The number of message returned is incorrect.")
		return
	}
	for i, message := range messages {
		formatted := formattedMessage[i]
		fmt.Printf(" * %s -> %s\n", message, formatted)

	}
}

func main() {
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addSignature)
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addGreeting)
}
