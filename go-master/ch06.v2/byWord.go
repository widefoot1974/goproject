package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Print(line)
	}

	return nil
}

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		// fmt.Print(line)

		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}

	return nil
}

func main() {

	fileName := "sample.txt"
	err := lineByLine(fileName)
	if err != nil {
		fmt.Printf("lineByLine(%v) fail: %v\n", fileName, err)
	}

	err = wordByWord(fileName)
	if err != nil {
		fmt.Printf("lineByLine(%v) fail: %v\n", fileName, err)
	}
}
