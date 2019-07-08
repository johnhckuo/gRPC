package main

import (
	"context"
	"log"
	"net"

	service "github.com/johnhckuo/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *service.SumRequest) (*service.SumResponse, error) {

	log.Printf("Executing sum operation...")
	firstInt := req.GetSum().GetFirstInt()
	secondInt := req.GetSum().GetSecondInt()
	sumRes := &service.SumResponse{
		Result: firstInt + secondInt,
	}
	return sumRes, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
