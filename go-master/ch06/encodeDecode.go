package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"created,omitempty"`
}

func main() {
	userall := UseAll{
		Name: "Mike",
		// Surname: "Tsoukalos",
		Year: 2021,
	}

	t, err := json.Marshal(&userall)
	if err != nil {
		fmt.Errorf("json.Marshal() fail: %v\n", err)
	} else {
		fmt.Printf("Value = %s\n", t)
		fmt.Printf("Value = %#v\n", t)
	}

	// str := `{"username":"Mike","surname":"Tsoukalos","created":2021}`
	// jsonRecord := []byte(str)

	temp := UseAll{}
	// err = json.Unmarshal(jsonRecord, &temp)
	err = json.Unmarshal(t, &temp)
	if err != nil {
		fmt.Errorf("json.Unmarshal() fail: %v\n", err)
	} else {
		fmt.Printf("Value = %T\n", temp)
		fmt.Printf("Value = %#v\n", temp)
	}
}
