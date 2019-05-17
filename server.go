package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From %q", html.EscapeString(r.URL.Path))
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", helloWorld)
	m.HandleFunc("/other/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Oter path")
	})
	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}
