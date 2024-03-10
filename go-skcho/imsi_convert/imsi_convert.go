package main

import (
	"encoding/base64"
	"fmt"
	"regexp"
)

type ImsiCoverter interface {
	ImsiEncode() string
	ImsiDecode(string) string
}

type Imsi struct {
	originImsi string
	encodeImsi string
}

func (imsi *Imsi) ImsiEncode() string {
	originalBytes := []byte(imsi.originImsi)
	imsi.encodeImsi = base64.StdEncoding.EncodeToString(originalBytes)

	fmt.Printf("originImsi = %d\n", imsi.originImsi)
	fmt.Printf("encodeImsi = %d\n", imsi.encodeImsi)
	return imsi.encodeImsi
}

func (imsi *Imsi) ImsiDecode() (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(imsi.encodeImsi)
	if err != nil {
		fmt.Printf("base64.StdEncoding.DecodeString(%s) fail: %v\n",
			imsi.encodeImsi, err)
		return "", err
	}
	imsi.originImsi = string(decodeBytes)

	fmt.Printf("encodeImsi = %v\n", imsi.encodeImsi)
	fmt.Printf("originImsi = %v\n", imsi.originImsi)
	return imsi.originImsi, nil
}

func extracetIMSI(subs_id string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(subs_id)
	if err != nil {
		fmt.Printf("base64.StdEncoding.DecodeString(%s) fail: %v\n", subs_id, err)
		return "", err
	}
	decodeString := string(decodeBytes)
	fmt.Printf("decodeString = %s\n", decodeString)

	re := regexp.MustCompile(`\d{15,16}`)
	matches := re.FindStringSubmatch(decodeString)
	if len(matches) > 0 {
		imsi := matches[0]
		if len(imsi) == 16 {
			return imsi[1:], nil
		} else if len(imsi) == 15 {
			return imsi, nil
		} else {
			return "", fmt.Errorf("Invalid Imsi format(%s)\n", imsi)
		}
	} else {
		return "", fmt.Errorf("Invalid Imsi format: decodeString(%s)", decodeString)
	}
}

func main() {

	subscriber_id := "AgAAOwEwNDUwMDUwMzc4MjM3NDEzQG5haS5lcGMubW5jMDA1Lm1jYzQ1MC4zZ3BwbmV0d29yay5vcmc="

	// imsi := Imsi{
	// 	encodeImsi: subscriber_id,
	// }

	// fmt.Printf(imsi.ImsiDecode())

	imsi, err := extracetIMSI(subscriber_id)
	if err != nil {
		fmt.Printf("extracetIMSI fail: %v\n", err)
	} else {
		fmt.Printf("imsi = %s\n", imsi)
	}

	fmt.Printf("")
}
