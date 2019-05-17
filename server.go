package main

import (
	"net/http"
)

func usersHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET method"))
	case http.MethodPost:
		w.Write([]byte("POST method"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Try Again. Method Not Allowed"))
	}
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/users", usersHandle)
	http.ListenAndServe(":8080", m)
}
