package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Brady-MacDonald/grpc/proto/hello" // Import generated code
	"google.golang.org/grpc"                         // gRPC client library
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server on localhost:50051
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	// Create a new HelloService client
	client := pb.NewHelloServiceClient(conn)

	// Set a timeout context for requests
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// REMOTE PROCEDURE CALL!
	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "World",
	})

	if err != nil {
		log.Fatal("Failed:", err)
	}

	log.Println("✅ Message received:", resp.Message)
}
