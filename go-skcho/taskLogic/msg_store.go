package main

type StoreData struct {
	// EIF 수신 메세지 데이터
	RecvHeader string
	Requests   []string
	Responses  []string

	// 유지 저장 데이터
	EIF_RECV_TOPIC  string
	AAAC_RECV_TOPIC string
	IMSI            string
	MSISDN          string
	Token           string

	// Task 데이터
	TaskCnt    int
	TaskingIdx int // 1부터 시작
	TaskStatus map[int]string
	cancel     string // timer
}
