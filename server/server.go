package main

import (
	"context"
	"log"
	"net"

	service "github.com/johnhckuo/gRPC/proto"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *service.GreetRequest) (*service.GreetResponse, error) {
	log.Printf("Greet function being invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &service.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Faield to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterDemoServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
