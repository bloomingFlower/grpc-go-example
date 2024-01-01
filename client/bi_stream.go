package main

import (
	"context"
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"io"
	"log"
	"time"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	// call the SayHelloBidirectionalStreaming method
	log.Printf("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("error while calling SayHelloBidirectionalStreaming RPC: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("error while reading stream: %v", err)
			}
			log.Printf("Response from SayHelloBidirectionalStreaming: %v", resp)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		helloRequest := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(helloRequest); err != nil {
			log.Fatalf("error while sending data to SayHelloBidirectionalStreaming RPC: %v", err)
		}
		log.Printf("Request sent to server: %s", name)
		time.Sleep(2 * time.Second)
	}
	err = stream.CloseSend()
	if err != nil {
		return
	}
	<-waitc
	log.Printf("Bidirectional Streaming Finished")
}
