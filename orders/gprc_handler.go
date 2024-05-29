package main

import (
	"context"
	pb "github.com/gauss2302/cafemania_commons/api"
	"google.golang.org/grpc"
	"log"
)

type GprcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &GprcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *GprcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Println("New order created")
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
