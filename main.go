package main

import (
	"context"
	"log"
	"net"

	"github.com/Brady-MacDonald/grpc/proto/hello" // Import generated code
	"google.golang.org/grpc"
)

// Implement the gRPC service
type helloServer struct {
	hello_proto.UnimplementedHelloServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *hello_proto.HelloRequest) (*hello_proto.HelloResponse, error) {
	return &hello_proto.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello_proto.RegisterHelloServiceServer(s, &helloServer{})

	x := hello_proto.HelloRequest{
		Name: "okay",
	}

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
