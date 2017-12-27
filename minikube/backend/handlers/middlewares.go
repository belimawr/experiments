package handlers

import (
	"log"
	"net/http"
	"time"
)

// NewLogMiddleware wraps a http.Handler with a logger
func NewLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Now().Sub(start)
		log.Printf("%s, duration: %.5fms",
			r.URL.String(),
			float64(elapsed)/float64(time.Millisecond))
	})
}
