package middleware

import (
	"log"
	"net/http"
)

func SetupLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		log.Printf("Request received: %s %s", r.Method, r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
