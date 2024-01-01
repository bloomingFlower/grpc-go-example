package main

import (
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received request from client: %v", req.Name)
		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + req.Name,
		}); err != nil {
			return err
		}
	}
}
