package util

import (
	"log"
	"net/http"
	"time"
)

func HttpLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		log.Printf("[%s] %s took %s", r.Method, r.URL.String(), elapsed)
	}
}
