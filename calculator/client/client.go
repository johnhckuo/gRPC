package main

import (
	"context"
	"log"

	service "github.com/johnhckuo/gRPC/calculator/proto"

	"google.golang.org/grpc"
)

func main() {
	serverURL := "localhost:50051"
	log.Println("Dialing %v", serverURL)
	cc, err := grpc.Dial(serverURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := service.NewCalculatorServiceClient(cc)
	doAdd(c, 3, 5)
}

func doAdd(c service.CalculatorServiceClient, param1 int64, param2 int64) {
	log.Println("Starting to do a Add operation")
	req := &service.SumRequest{
		Sum: &service.Sum{
			FirstInt:  param1,
			SecondInt: param2,
		},
	}
	result, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling add rpc: %v", err)
	}
	log.Printf("Response from calculator: %v\n", result)

}
