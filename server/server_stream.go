package main

import (
	"log"
	"time"

	pb "github.com/ak-yudha/grpc-stream/proto"
)

func (s *helloServer) ProductServerStreaming(req *pb.NamesList, stream pb.GreetService_ProductServerStreamingServer) error {
	log.Printf("Got request with names : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.ProductResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
