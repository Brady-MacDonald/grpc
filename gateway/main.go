package main

import (
	"log"

	"github.com/Brady-MacDonald/grpc/gateway/clients"
	"github.com/Brady-MacDonald/grpc/gateway/config"
	"github.com/Brady-MacDonald/grpc/gateway/server"
)

func main() {
	cfg := config.Load()

	log.Printf(
		"starting gateway on %s (shortener: %s)",
		cfg.HTTPAddr,
		cfg.ShortenerGRPCAddr,
	)

	_ = clients.New(cfg)

	server := server.New(cfg)
	server.Start()
}
