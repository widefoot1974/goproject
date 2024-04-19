package main

import (
	"fmt"
)

type StoreData struct {
	msisdn      string
	totalReqCnt int
	taskStatus  map[int]string
	sessStatus  []int
}

func task1(tStatus map[int]string) {
	tStatus[0] = "Zero"
	tStatus[1] = "One"
	tStatus[2] = "Two"
}

func task2(sessStatus *[]int) {
	*sessStatus = append(*sessStatus, 6000)
	*sessStatus = append(*sessStatus, 6001)
	*sessStatus = append(*sessStatus, 6002)

	fmt.Printf("sessStatus = %#v\n", sessStatus)
}

func main() {

	storeData := StoreData{
		msisdn:      "01044123559",
		totalReqCnt: 3,
		taskStatus:  map[int]string{},
		sessStatus:  []int{},
	}

	taskStatus := storeData.taskStatus
	sessStatus := &storeData.sessStatus

	fmt.Printf("storeData = %#v\n", storeData)

	task1(taskStatus)

	fmt.Printf("storeData = %#v\n", storeData)

	task2(sessStatus)

	fmt.Printf("storeData = %#v\n", storeData)

	fmt.Printf("")
}
