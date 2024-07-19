package main

import (
	"log"
	"strings"
)

func main() {

	var queryString string = "IMSI_EAP=AgAAOwEwNDUwMDUwMjA5OTIwMDAwQG5haS5lcGMubW5jMDA1Lm1jYzQ1MC4zZ3BwbmV0d29yay5vcmc=&terminal_id=358333209920000"
	imsi := getIMSIFromQueryString(queryString)

	log.Printf("imsi = [%v]\n", imsi)
}

func getIMSIFromQueryString(queryString string) string {
	log.Printf("queryString = [%v]\n", queryString)
	params := strings.Split(queryString, "&")
	var eapId, imsiEap string

	log.Printf("params = %v\n", params)

	for _, param := range params {
		log.Printf("param = %v\n", param)
		keyValue := strings.SplitN(param, "=", 2)
		if len(keyValue) != 2 {
			log.Printf("len(keyValue) = %v\n", len(keyValue))
			continue
		}
		key, value := keyValue[0], keyValue[1]
		log.Printf("param key = [%v], value = [%v]\n", key, value)
		if key == "EAP_ID" {
			eapId = value
		} else if key == "IMSI_EAP" {
			imsiEap = value
		}
	}

	log.Printf("eapId = [%v]\n", eapId)
	log.Printf("imsiEap = [%v]\n", imsiEap)

	if eapId != "" {
		return eapId
	}

	if imsiEap != "" {
		return imsiEap
	}

	return ""
}
