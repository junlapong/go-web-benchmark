package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := http.NewServeMux()
	router.HandleFunc("/", hello)
	router.HandleFunc("/ping", ping)
	r := loggerMiddleware(router)

	log.Println("Start http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hello, World!")
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"pong"}`)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			lrw := newLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)
			log.Printf("%d\t%s\t%s", lrw.StatusCode, r.Method, r.URL.Path)
		})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
