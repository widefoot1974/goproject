package main

import (
	"context"
	"log"
	"net"
	"youtube/MaximilienAndile/demo-grpc/invoicer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

// func main() {
// 	lis, err := net.Listen("tcp", ":8089")
// 	if err != nil {
// 		log.Fatalf("can't create listener: %v\n", err)
// 	}
// 	serverRegistrar := grpc.NewServer()
// 	service := &myInvoicerServer{}
// 	invoicer.RegisterInvoicerServer(serverRegistrar, service)

// 	log.Printf("server starting...\n")

// 	err = serverRegistrar.Serve(lis)
// 	if err != nil {
// 		log.Fatalf("impossible to serve: %v\n", err)
// 	}
// }

func main() {
	if err := runServer(); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func runServer() error {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		return err
	}
	defer lis.Close()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logInterceptor),
	}
	serverRegistrar := grpc.NewServer(opts...)
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	log.Printf("Serve is listening...\n")

	return serverRegistrar.Serve(lis)
}

// logInterceptor는 각 요청에 대한 로깅 인터셉터입니다.
func logInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Received request from %s - Method: %s\n", peer.Addr, info.FullMethod)
	}
	return handler(ctx, req)
}
