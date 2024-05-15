package messageratemonitor

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type RateMonitor struct {
	data           *ThreadData
	wg             sync.WaitGroup
	processFunc    func(int) int
	ticker         *time.Ticker
	stopTickerChan chan bool
}

type ThreadData struct {
	Mutex     sync.Mutex
	Sum       int64
	Max       int
	ItemCount int64
}

func NewRateMonitor(processFunc func(int) int) *RateMonitor {
	return &RateMonitor{
		data:           &ThreadData{},
		processFunc:    processFunc,
		stopTickerChan: make(chan bool),
	}
}

func (rm *RateMonitor) Add(value int) {
	atomic.AddInt64(&rm.data.Sum, int64(value))
	atomic.AddInt64(&rm.data.ItemCount, 1)

	rm.data.Mutex.Lock()
	if value > rm.data.Max {
		rm.data.Max = value
	}
	rm.data.Mutex.Unlock()
}

func (rm *RateMonitor) StartMonitoring(numGoroutines, loopCount int, sleepDuration, reportInterval time.Duration) {
	for i := 0; i < numGoroutines; i++ {
		rm.wg.Add(1)
		go func() {
			defer rm.wg.Done()
			for j := 0; j < loopCount; j++ {
				time.Sleep(sleepDuration)
				value := rand.Intn(1000)
				processedValue := rm.processFunc(value)
				rm.Add(processedValue)
			}
		}()
	}

	if reportInterval > 0 {
		rm.ticker = time.NewTicker(reportInterval)
		go rm.reportPeriodically()
	}
}

func (rm *RateMonitor) reportPeriodically() {
	for {
		select {
		case <-rm.ticker.C:
			sum, max, itemCount := rm.getCurrentStats()
			log.Printf("Periodic Report - Total Sum: %d, Max Value: %d, Total Items Processed: %d\n", sum, max, itemCount)
			rm.Reset() // Reset data after each report
		case <-rm.stopTickerChan:
			return
		}
	}
}

func (rm *RateMonitor) getCurrentStats() (int64, int, int64) {
	rm.data.Mutex.Lock()
	sum := atomic.LoadInt64(&rm.data.Sum)
	max := rm.data.Max
	itemCount := atomic.LoadInt64(&rm.data.ItemCount)
	rm.data.Mutex.Unlock()
	return sum, max, itemCount
}

func (rm *RateMonitor) WaitForCompletion() (int64, int, int64) {
	rm.wg.Wait()
	if rm.ticker != nil {
		rm.ticker.Stop()
		rm.stopTickerChan <- true
	}
	return rm.getCurrentStats()
}

func (rm *RateMonitor) Reset() {
	atomic.StoreInt64(&rm.data.Sum, 0)
	atomic.StoreInt64(&rm.data.ItemCount, 0)
	rm.data.Mutex.Lock()
	rm.data.Max = 0
	rm.data.Mutex.Unlock()
}

