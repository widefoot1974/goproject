package main

import (
	"log"
	"net"
	"youtube/MaximilienAndile/demo-grpc/invoicer"
)

func main() {

	lis, err := net.Listen("tcp", ":8089:")
	if err != nil {
		log.Fatalf("can't creater listener: %v\n", err)
	}

	invoicer.RegisterInvoicerServer()
}
