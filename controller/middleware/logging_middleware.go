package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleWare(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(rw, req)
		log.Printf("Uri: %s | Method: %s | Time Interval: %s", req.RequestURI, req.Method, time.Since(startTime))
	}
}
