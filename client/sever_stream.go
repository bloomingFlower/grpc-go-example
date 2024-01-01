package main

import (
	"context"
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("error while calling SayHelloServerStreaming RPC: %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from SayHelloServerStreaming: %v", resp)
	}
	log.Printf("Streaming Finished")
}
