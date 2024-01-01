package main

import (
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received request from client: %v", req.Names)
	for _, name := range req.Names {
		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + name,
		}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
