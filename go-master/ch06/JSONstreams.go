package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		fmt.Printf("i = %#v, myRand = %#v, newChar = %#v\n", i, myRand, newChar)
		if i == l {
			break
		}
		i++
	}
	return temp
}

func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func main() {
	var t Data
	for i := 0; i < 2; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		fmt.Printf("t = %#v\n", t)
		DataRecords = append(DataRecords, t)
	}

	buf1 := new(bytes.Buffer)
	encoder := json.NewEncoder(buf1)
	encoder.SetIndent("", "  ")
	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Printf("Serialize() fail: %v\n", err)
		return
	}

	fmt.Printf("Type buf: %T\n", buf1)
	fmt.Printf("After Serialize: %v\n", buf1)
	fmt.Printf("After Serialize: %#v\n", buf1)

	buf2 := new(bytes.Buffer)
	decoder := json.NewDecoder(buf2)
	var temp []Data
	err = DeSerialize(decoder, &temp)
	if err != nil {
		fmt.Printf("DeSerialize() fail: %v\n", err)
	}

	fmt.Printf("After DeSerialize: %#v\n", temp)
}
