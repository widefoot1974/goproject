package main

import (
	"context"
	"fmt"
	"go-master/ch12/protoapi"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, r *protoapi.RequestDateTime) (*protoapi.DateTime, error) {
	currentTime := time.Now()
	response := &protoapi.DateTime{
		Value: currentTime.String(),
	}
	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *protoapi.RandomParams) (*protoapi.RandomInt, error) {
	rand.Seed(r.GetSeed())
	place := r.GetPlace()

	min, max := 0, 100 // 예시 범위
	temp := rand.Intn(max-min+1) + min
	for {
		place--
		if place <= 0 {
			break
		}
		temp = rand.Intn(max-min+1) + min
	}

	response := &protoapi.RandomInt{
		Value: int64(temp),
	}

	return response, nil
}

func (RandomServer) GetRandomPass(ctx context.Context, r *protoapi.RequestPass) (*protoapi.RandomPass, error) {
	rand.Seed(r.GetSeed())
	temp := generateRandomString(r.GetLength())

	response := &protoapi.RandomPass{
		Password: temp,
	}
	return response, nil
}

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()

	var randomServer RandomServer
	protoapi.RegisterRandomServer(server, randomServer)
	reflection.Register(server)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server requesets...")
	server.Serve(listen)
}

// generateRandomString 함수는 주어진 길이의 랜덤 문자열을 생성합니다.
func generateRandomString(length int64) string {
	// 사용할 문자 집합을 정의합니다.
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // 랜덤 시드 설정
	b := make([]byte, length)                                     // 문자열 길이에 맞는 바이트 슬라이스 생성

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))] // 문자 집합에서 랜덤하게 문자를 선택하여 할당
	}
	return string(b) // 바이트 슬라이스를 문자열로 변환하여 반환
}
