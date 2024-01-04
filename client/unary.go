package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ak-yudha/grpc-stream/proto"
)

// callProduct will need to be called from main.go to make the call to the server
func callProduct(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Product(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
