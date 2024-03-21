package main

import (
	"fmt"
)

func handle_eif_mag(taskCh chan<- int) {

	for {
		// nats-server로부터 메세지 수신

		// 메세지로부터 header, content 추출

		// storeId를 생성하고, MsgStore에 저장

		// taskCh을 통해서 msg_task()로 전달

		fmt.Printf("")
	}
}
