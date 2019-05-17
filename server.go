package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	server := &http.Server{

		Addr:           ":8080",
		Handler:        m,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server listen at :8080")
	log.Fatal(server.ListenAndServe())
}
