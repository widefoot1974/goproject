package main

import (
	handler "kitchen/services/orders/handler/orders"
	"kitchen/services/orders/service"
	"log"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	ordersService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(ordersService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
