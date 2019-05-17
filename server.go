package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	log.Println("Server listen at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
