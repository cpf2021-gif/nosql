package main

import (
	"log"
	"net/http"
)

// go run cmd/server/server.go
func main() {
	fs := http.FileServer(http.Dir("static"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
