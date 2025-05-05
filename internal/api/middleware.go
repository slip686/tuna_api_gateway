package api

import (
	"fmt"
	"net/http"
	"time"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseWriter := NewLoggingResponseWriter(w)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(responseWriter, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(responseWriter, r)
	})
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseWriter := NewLoggingResponseWriter(w)
		next.ServeHTTP(responseWriter, r)
		fmt.Printf("[%s] %s, %s %d %s\n", r.Method, r.RemoteAddr, r.URL.Path,
			responseWriter.statusCode, time.Since(start))
	})
}
