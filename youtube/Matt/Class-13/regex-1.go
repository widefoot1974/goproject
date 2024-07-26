package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func B() string {
	_, file, line, _ := runtime.Caller(1)
	idx := strings.LastIndex(file, "/")
	fileName := file[idx+1:]
	fmt.Printf("file = %v, line = %v\n", file, line)
	return fileName + ":" + strconv.Itoa(line)
}

func A() string {
	return B()
}

func main() {

	fmt.Printf("")
	fmt.Println(A())
}
