package main

import (
	"encoding/xml"
	"fmt"
)

type Employee struct {
	XMLName   xml.Name `xml:"employee"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
	Address
	Comment string `xml:",comment"`
}

type Address struct {
	City, Country string
}

func main() {
	r := Employee{ID: 7, FirstName: "Muhalis", LastName: "Tsoukalos"}
	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	output, err := xml.MarshalIndent(&r, "", "    ")
	if err != nil {
		fmt.Printf("xml.MarshalIndent() fail: %v\n", err)
	}

	output = []byte(xml.Header + string(output))
	fmt.Printf("output = \n%s\n", output)
	fmt.Printf("output = \n%v\n", output)
}
