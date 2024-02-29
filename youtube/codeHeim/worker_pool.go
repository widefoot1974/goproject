package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Way to process the tasks
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %v\n", t.Email)
	// Simulate a time consuming process
	time.Sleep(time.Second * 2)
}

// Email task definition
type ImageProccessingTask struct {
	ImageUrl string
}

// Way to process the tasks
func (t *ImageProccessingTask) Process() {
	fmt.Printf("Processing the image %v\n", t.ImageUrl)
	// Simulate a time consuming process
	time.Sleep(time.Second * 5)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// Start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	// Wait for all tasks finish
	wp.wg.Wait()
}
