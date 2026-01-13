package main

import (
	"log"
	"net"
	"os"

	shortenerv1 "github.com/Brady-MacDonald/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	store := NewStore()
	server := NewServer(store)

	addr := ":50051"
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server running on :50051")

	grpcServer := grpc.NewServer()
	shortenerv1.RegisterShortenerServiceServer(grpcServer, server)

	log.Println("Shortener service listening on", addr)
	log.Fatal(grpcServer.Serve(lis))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
