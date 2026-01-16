package server

import (
	"github.com/Brady-MacDonald/grpc/gateway/handlers"
)

// NewRouter creates a new Router instance
// Map routes to handlers
func InjectRoutes(s *Server) {
	/** Create Services (Dependencies) **/
	// cmsService := cmsserivce.New(cmsConfig)

	/** Create Handlers Providing any Dependencies **/
	shortenHandler := handlers.NewShortener()
	redirectsHandler := handlers.NewRedirects()

	/** Middleware **/
	// commLogger := middleware.LogReq("commercial_content", commercialHandler.Id)

	/** Routes **/
	s.server.Handle("POST /shorten", logMiddleware(shortenHandler.ShortenHandler))
	s.server.HandleFunc("GET /{slug}", logMiddleware(redirectsHandler.RedirectHandler))
}
