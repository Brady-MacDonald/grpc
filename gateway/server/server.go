package server

import (
	"fmt"
	"net/http"

	"github.com/Brady-MacDonald/grpc/gateway/config"
)

type Server struct {
	port   int
	server *http.ServeMux
}

// Creates a new Server instance
func New(cfg *config.Config) *Server {
	s := &Server{
		port:   cfg.Port,
		server: http.NewServeMux(),
	}

	InjectRoutes(s)
	return s
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
	fmt.Printf("Server starting on port %d\n", addr)

	return http.ListenAndServe(addr, s.server)
}
