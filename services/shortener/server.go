package main

import (
	"context"
	"errors"
	pb "github.com/Brady-MacDonald/grpc/proto"
)

type server struct {
	store *Store
	pb.UnimplementedShortenerServiceServer
}

// Create new Server
// Implements the proto contract
func NewServer(store *Store) *server {
	return &server{
		store: store,
	}
}

func (s *server) CreateShortenedUrl(ctx context.Context, req *pb.CreateShortenedUrlRequest) (*pb.CreateShortenedUrlResponse, error) {
	if req.OriginalUrl == "" {
		return nil, errors.New("original URL is required")
	}

	resp := &pb.CreateShortenedUrlResponse{
		ShortenedUrl: "",
	}

	return resp, nil
}

func (s *server) ResolveURL(ctx context.Context, req *pb.ResolveShortenedUrlRequest) (*pb.ResolveShortenedUrlResponse, error) {
	if req.ShortenedUrl == "" {
		return nil, errors.New("shortened URL is required")
	}

	resp := &pb.ResolveShortenedUrlResponse{
		OriginalUrl: "",
	}

	return resp, nil
}
