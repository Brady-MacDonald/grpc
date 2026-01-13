package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Brady-MacDonald/grpc/proto/hello" // Import generated code
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

// Implement the service defined in proto buf
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := &pb.HelloResponse{
		Message: fmt.Sprintf("Hello there %s!", req.Name),
	}

	return resp, nil
}

func main() {
	// Listen on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server running on :50051")

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
