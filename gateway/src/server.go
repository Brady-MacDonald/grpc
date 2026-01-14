package gateway

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
}

// Create new server
func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func Hi(w http.ResponseWriter, r *http.Request) {
	h := r.Host
	res := fmt.Sprintf("Hey: %s", h)
	fmt.Fprintf(w, "%s", res)
}

// Handle the POST request
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

}

func (s *Server) Listen() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", Hi)
	mux.HandleFunc("/shorten", ShortenHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), mux)

	return err
}
