package main

import (
	"log"
	"net/http"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] [%d] %s", r.Method, 200, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func main() {
	m := http.NewServeMux()
	home := http.HandlerFunc(homePage)
	m.Handle("/", logMiddleware(home))
	log.Println("Server listen at :8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}
