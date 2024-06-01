package main

import (
	"fmt"
	"go-master/ch12/protoapi"
)

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

func main() {
	fmt.Println("vim-go")
}
