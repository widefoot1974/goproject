package main

import (
	"context"
	"log"
	"time"
	"youtube/MaximilienAndile/demo-grpc/invoicer"

	"google.golang.org/grpc"
)

func main() {
	conn, err := connectGRPC("localhost:8089")
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := invoicer.NewInvoicerClient(conn)
	createInvoice(client)
}

func connectGRPC(address string) (*grpc.ClientConn, error) {
	// 연결을 시도하기 전에 로그를 출력합니다.
	log.Printf("Attempting to connect to gRPC server at %s", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// 연결 성공 시 로그를 남깁니다.
	log.Printf("Successfully connected to %s", address)

	return conn, nil
}

// createInvoice sends a Create request to the Invoicer service and logs the response.
func createInvoice(client invoicer.InvoicerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:   10000,
			Currency: "USD",
		},
		From: "Sender Name",
		To:   "Receiver Name",
	}

	resp, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create invoice: %v", err)
	}

	log.Printf("Received PDF length: %d, value:%v", len(resp.Pdf), string(resp.Pdf))
	log.Printf("Received DOCX length: %d, value:%v", len(resp.Docx), string(resp.Docx))
}
