package main

import (
	"fmt"
)

func main() {
	// Create new tasks
	// tasks := make([]Task, 20)
	// for i := 0; i < 20; i++ {
	// 	tasks[i] = Task{ID: i + 1}
	// }

	tasks := []Task{
		&EmailTask{Email: "widefoot1111@naver.com", Subject: "travel", MessageBody: "~~~"},
		&ImageProccessingTask{ImageUrl: "/images/sample1.jpg"},
		&EmailTask{Email: "widefoot2222@naver.com", Subject: "travel", MessageBody: "~~~"},
		&ImageProccessingTask{ImageUrl: "/images/sample2.jpg"},
		&EmailTask{Email: "widefoot3333@naver.com", Subject: "travel", MessageBody: "~~~"},
		&ImageProccessingTask{ImageUrl: "/images/sample3.jpg"},
		&EmailTask{Email: "widefoot4444@naver.com", Subject: "travel", MessageBody: "~~~"},
		&ImageProccessingTask{ImageUrl: "/images/sample4.jpg"},
		&EmailTask{Email: "widefoot5555@naver.com", Subject: "travel", MessageBody: "~~~"},
		&ImageProccessingTask{ImageUrl: "/images/sample5.jpg"},
		&EmailTask{Email: "widefoot6666@naver.com", Subject: "travel", MessageBody: "~~~"},
		&ImageProccessingTask{ImageUrl: "/images/sample6.jpg"},
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // Number of workers that can run at a time
	}

	// Run the pool
	wp.Run()

	fmt.Println("All Tasks have been processed!")
}
