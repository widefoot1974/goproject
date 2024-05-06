package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func main() {
	// Create a new Person
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "Johndoe@gmail.com",
	}

	// Write the Person to an XML file
	file, err := os.Create("person.xml")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// 구조체를 XML로 변환
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	err = encoder.Encode(person)
	if err != nil {
		log.Panic(err)
	}

	xmlData, err := xml.MarshalIndent(person, "", "  ")
	// xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Panicf("[ERROR] xml.MarshalIndent() Fail: %v\n", err)
		return
	}
	log.Printf("xmlData Type is %T\n", xmlData)
	log.Printf("xmlData size is %v", len(xmlData))
	log.Printf("xmlData =\n%v\n", string(xmlData))

	xmlInput := string(xmlData)
	var p2 Person
	err = xml.Unmarshal([]byte(xmlInput), &p2)
	if err != nil {
		log.Panicf("[ERROR] xml.Unmarshal() Fail: %v\n", err)
		return
	}
	fmt.Printf("p2 = %+v\n", p2)

	// Read the Person from the XML file
	file, err = os.Open("person.xml")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	var newPerson Person
	err = decoder.Decode(&newPerson)
	if err != nil {
		log.Panic(err)
	}

	// Print the new Person
	fmt.Println("New Person:")
	fmt.Println("Name", newPerson.Name)
	fmt.Println("Age", newPerson.Age)
	fmt.Println("Email", newPerson.Email)
}
