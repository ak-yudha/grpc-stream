package main

import (
	"io"
	"log"

	pb "github.com/ak-yudha/grpc-stream/proto"
)

func (s *helloServer) ProductClientStreaming(stream pb.GreetService_ProductClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
