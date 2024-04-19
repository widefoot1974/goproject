package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	logFile    *os.File
	lastLogMin int
)

func SetLog() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	now := time.Now()
	updateLogFile(now)
	lastLogMin = now.Minute()

	for {
		select {
		case now := <-ticker.C:
			// 매 분의 정각 00초에만 로그 파일을 업데이트
			if lastLogMin != now.Minute() {
				updateLogFile(now)
				lastLogMin = now.Minute()
			}
		}
	}
}

func updateLogFile(now time.Time) {
	currentDateTime := now.Format("2006-01-02_15-04")
	logFileName := fmt.Sprintf("%s.%s.log", "ios", currentDateTime)
	if logFile != nil && lastLogMin == time.Now().Minute() {
		return // 같은 분에 다시 파일을 열지 않도록 함
	}

	CloseLog() // 이전 로그 파일 닫기

	var err error
	logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("File(%v) Open Fail: %v\n", logFileName, err)
		return
	}

	lastLogMin = time.Now().Minute()

	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	multi := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multi)
}

func CloseLog() {
	if logFile != nil {
		logFile.Close()
		logFile = nil // 파일을 닫은 후 nil로 설정하여 다음 로그 파일을 열 준비
	}
}

func main() {
	go SetLog() // SetLog를 고루틴으로 실행하여 독립적으로 로그 관리

	// 메인 로직 처리 (예시)
	for {
		log.Printf("now = %v\n", time.Now())
		time.Sleep(5 * time.Second) // 메인 작업 로직, 여기서는 단순히 대기
	}
}
