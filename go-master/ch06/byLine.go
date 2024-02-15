package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	fmt.Println("--------------------------------------------------")
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading file %s", err)
		}
		fmt.Print(line)
	}
	fmt.Println("--------------------------------------------------")

	return nil
}

func main() {

	file := "test.tx"
	err := lineByLine(file)
	if err != nil {
		fmt.Println("lineByLine fail:", err)
	}
}
