package main

import (
	"fmt"
)

func main() {

	fmt.Printf("")

	taskCh := make(chan int)

	go handle_eif_msg(taskCh)

	go handle_aaac_msg(taskCh)

	go msg_task(taskCh)

}
