package main

import (
	"context"
	"log"

	service "github.com/johnhckuo/gRPC/proto"
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
	c := service.NewDemoServiceClient(cc)
	doUnary(c)
}

func doUnary(c service.DemoServiceClient) {
	log.Println("Starting to do a Unary RPC...")
	req := &service.GreetRequest{
		Greeting: &service.Greeting{
			FirstName: "john",
			LastName:  "Kuo",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc: %v", err)
	}
	log.Printf("Response from Greet: %v\n", res.Result)
}
