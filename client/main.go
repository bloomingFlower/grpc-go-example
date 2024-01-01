package main

import (
	pb "github.com/bloomingFlower/grpc-go-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	// make a connection to the server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error closing the connection: %v", err)
		}
	}(conn)

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"John", "Doe", "Jane"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}
