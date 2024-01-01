package main

import (
	"context"
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	// call the SayHello method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}
	log.Printf("Response from server: %s", res.Message)
}
