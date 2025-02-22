package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1   S1
	text []byte
}

func (s *S1) Read(p []byte) (int, error) {
	fmt.Printf("p = %T, %#v\n", p, p)
	fmt.Print("Give me your name: ")
	fmt.Scanln(&p)
	fmt.Printf("p = %T, %#v\n", p, p)
	s.F2 = string(p)
	return len(p), nil
}

func (s *S1) Write(p []byte) (int, error) {
	fmt.Printf("s.F1 = %v\n", s.F1)
	if s.F1 < 0 {
		return -1, nil
	}

	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s ", p)
	}
	fmt.Println()
	return s.F1, nil
}

func (s S2) eof() bool {
	return len(s.text) == 0
}

func (s *S2) readByte() byte {
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return
	}

	l := len(p)
	if l > 0 {
		for n < 1 {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return
}

func main() {

	s1var := S1{4, "Hello"}
	fmt.Printf("s1var = %v\n", s1var)

	buf := make([]byte, 2)
	n, err := s1var.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("n = %v\n", n)
	fmt.Printf("Read: %v\n", s1var.F2)
	_, _ = s1var.Write([]byte("Hello There!"))

	s2var := S2{F1: s1var, text: []byte("Hello world!!")}

	r := bufio.NewReader(&s2var)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("-", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}
}
