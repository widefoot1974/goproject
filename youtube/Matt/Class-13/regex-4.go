package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// UUID validation
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter UUID: ")
	uuidInput, _ := reader.ReadString('\n')
	uuidInput = uuidInput[:len(uuidInput)-1] // Remove newline character

	err := checkUUID(uuidInput)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UUID is valid.")
	}
}

func checkUUID(ustr string) error {

	uure := `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`
	ufmt := regexp.MustCompile(uure)

	if !ufmt.MatchString(ustr) {
		return fmt.Errorf("%s is not a UUID", ustr)
	}

	return nil
}
