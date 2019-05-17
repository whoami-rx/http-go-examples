package main

import (
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	m.Handle("/", fs)
	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}
