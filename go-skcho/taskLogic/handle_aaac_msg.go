package main

import (
	"fmt"
)

func handle_aaac_mag(taskCh chan<- int) {

	for {
		// nats-server로부터 메세지 수신

		// 메세지로부터 header, content 추출

		// header로부터 storeId를 추출

		// storeId로 메세지를 조회해서, Responsee 메세지에 결과 저장

		// taskCh을 통해서 msg_task()로 전달
		taskCh <- storeId

		fmt.Printf("")
	}

}
