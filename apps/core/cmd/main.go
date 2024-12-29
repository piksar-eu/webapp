package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	mux := http.NewServeMux()

	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})

	log.Printf("Serve app on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)

	if err != nil {
		panic(err)
	}
}
