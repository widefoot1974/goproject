package main

import "log"

func msg_task(taskCh <-chan int) {

	for storeId := range taskCh {

		log.Printf("received storeId = %v\n", storeId)

		// storeId로 StoreData 정보를 가져옴

		// 모든 Request 처리가 끝난 메세지는 EIF로 결과 전송

	}

}
