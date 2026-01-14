package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Brady-MacDonald/grpc/gateway/handlers"
)

func main() {
	httpAddr := getEnv("HTTP_ADDR", ":8080")
	// shortenerGRPCAddr := getEnv("SHORTENER_GRPC_ADDR", "localhost:50051")

	// --- gRPC connection ---
	// conn, err := grpc.Dial(
	// 	shortenerGRPCAddr,
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalf("failed to connect to shortener service: %v", err)
	// }
	// defer conn.Close()
	//
	// shortenerClient := grpcclient.NewShortenerClient(conn)

	// --- HTTP handlers ---
	handler := handlers.NewHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/", handler.Redirect)

	server := &http.Server{
		Addr:         httpAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// --- graceful shutdown ---
	go func() {
		log.Printf("API Gateway listening on %s", httpAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %v", err)
		}
	}()

	// waitForShutdown(server)
}

// Helper func to return ENV variables
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
