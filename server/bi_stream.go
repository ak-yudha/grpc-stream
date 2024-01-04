package main

import (
	"io"
	"log"

	pb "github.com/ak-yudha/grpc-stream/proto"
)

func (s *helloServer) ProductBidirectionalStreaming(stream pb.GreetService_ProductBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.ProductResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
