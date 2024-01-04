package main

import (
	"context"

	pb "github.com/ak-yudha/grpc-stream/proto"
)

func (s *helloServer) Product(ctx context.Context, req *pb.NoParam) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{
		Message: "Hello unary",
	}, nil
}
