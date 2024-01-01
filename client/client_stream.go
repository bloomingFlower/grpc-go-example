package main

import (
	"context"
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	// call the SayHelloClientStreaming method
	log.Printf("Client Streaming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("error while calling SayHelloClientStreaming RPC: %v", err)
	}
	for _, name := range names.Names {
		helloRequest := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(helloRequest); err != nil {
			log.Fatalf("error while sending data to SayHelloClientStreaming RPC: %v", err)
		}
		log.Printf("Request sent to server: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from SayHelloClientStreaming RPC: %v", err)
	}
	log.Printf("Response from SayHelloClientStreaming: %v", res)
	log.Printf("Client Streaming Finished")
}
