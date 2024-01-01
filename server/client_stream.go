package main

import (
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Received request from client: %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
