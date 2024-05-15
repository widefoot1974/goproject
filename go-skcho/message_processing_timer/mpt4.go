package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type ThreadData struct {
	Mutex     sync.Mutex
	Sum       int
	Max       int
	ItemCount int
}

func (td *ThreadData) Add(value int) {
	td.Mutex.Lock()
	defer td.Mutex.Unlock()
	if value > td.Max {
		td.Max = value
	}
	td.Sum += value
	td.ItemCount += 1
}

var (
	sharedData *ThreadData
	wg         sync.WaitGroup
)

func init() {
	sharedData = &ThreadData{}
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	thrCnt := 5
	for i := 0; i < thrCnt; i++ {
		wg.Add(1)
		go task()
	}

	// 매초 통계 계산
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			computeOverallStats()
		}
	}()

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다립니다.
	ticker.Stop()
	computeOverallStats() // 최종 계산
}

func task() {
	defer wg.Done()
	loopCnt := 100000
	for i := 0; i < loopCnt; i++ {
		time.Sleep(time.Millisecond * 5)
		randNum := rand.Intn(1000)
		sharedData.Add(randNum)
	}
}

func computeOverallStats() {
	sharedData.Mutex.Lock()
	totalCount := sharedData.ItemCount
	overallSum := sharedData.Sum
	overallMax := sharedData.Max

	// 값을 초기화합니다.
	sharedData.Sum = 0
	sharedData.ItemCount = 0
	sharedData.Max = 0
	sharedData.Mutex.Unlock()

	average := float64(overallSum) / float64(totalCount)
	log.Printf("Total Values Count: %d, Current Overall Average: %.2f, Current Overall Max: %d\n", totalCount, average, overallMax)
}
