package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port              int
	HTTPAddr          string
	ShortenerGRPCAddr string
	RequestTimeout    time.Duration
}

func Load() *Config {
	cfg := &Config{
		HTTPAddr:          getEnv("HTTP_ADDR", ":8080"),
		ShortenerGRPCAddr: getEnv("SHORTENER_GRPC_ADDR", "localhost:50051"),
		RequestTimeout:    getDurationEnv("REQUEST_TIMEOUT_MS", 200),
	}

	validate(cfg)
	return cfg
}

/**********************
*** PRIVATE HELPERS ***
***********************/

func validate(cfg *Config) {
	if cfg.HTTPAddr == "" {
		log.Fatal("HTTP_ADDR cannot be empty")
	}

	if cfg.ShortenerGRPCAddr == "" {
		log.Fatal("SHORTENER_GRPC_ADDR cannot be empty")
	}

	if cfg.RequestTimeout <= 0 {
		log.Fatal("REQUEST_TIMEOUT_MS must be greater than zero")
	}
}

// Retrieves the value of the environment variable named by the key
// Otherwise provides fallback value
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}

func getDurationEnv(key string, fallbackMs int) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return time.Duration(fallbackMs) * time.Millisecond
	}

	ms, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("invalid value for %s: %v", key, err)
	}

	return time.Duration(ms) * time.Millisecond
}
