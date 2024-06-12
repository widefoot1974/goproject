package handler

import (
	"context"
	"kitchen/services/common/genproto/orders"
	"kitchen/services/orders/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService() {
	gRPCHandler := &OrdersGrpcHandler{}

	// register the OrderServiceServer
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
