package main

import (
	"log"

	pb "github.com/Go_Projects/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"Suryansh", "Krishna", "Saurabh", "Abhishek"},
	// }
	callSayHello(client)
}
