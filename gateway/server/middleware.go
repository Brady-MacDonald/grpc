package server

import (
	"log"
	"net/http"
	"time"
)

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// Return a HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the method and the requested URL
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log how long it took
		log.Printf("Completed in %v", time.Since(start))
	}
}

func errorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error and send a user-friendly message
				log.Printf("Error occurred: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
