package main

import (
	"fmt"
)

type Message struct {
	MsgId    int32
	Name     string
	contents []byte
}

func main() {

	fmt.Printf("")

	ptrMsg := NewMessage()
	fmt.Printf("msg = %#v\n", ptrMsg)

	// copy pointer data to struct
	msg := *ptrMsg
	fmt.Printf("msg = %#v\n", msg)
}

func NewMessage() *Message {
	return &Message{
		MsgId:    1,
		Name:     "first",
		contents: []byte("first"),
	}
}
