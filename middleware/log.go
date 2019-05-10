package middleware

import (
	"log"
	"net/http"
)

// Log requests
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// FIXME: prettier & more informative
		log.Printf("%s %s", r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
