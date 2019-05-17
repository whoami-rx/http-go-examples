package main

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler custom http handler
type ErrorHandler struct {
	Message string
	Status  int
}

func (e ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(e.Status)
	fmt.Fprintf(w, "Message: %s\nStatus: %d", e.Message, e.Status)
}

func main() {
	e := ErrorHandler{Message: "Custom NOT FOUND ERROR", Status: 404}
	m := http.NewServeMux()
	m.Handle("/", e)
	log.Println("Server listen at :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}
