package main

import (
	"io"
	"log"

	pb "github.com/Go_Projects/gRPC/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got rqeuest with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
