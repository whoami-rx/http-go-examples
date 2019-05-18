package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User ...
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JSONMiddleware only set content type: application/json
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(User{
			Email:    "example@user.com",
			Password: "NotSendPasswords",
		})
	})
	m.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			fmt.Fprintf(w, "%q", user)
		default:
			w.Write([]byte("Try with post"))
		}
	})
	log.Println("Server listen at :8080")
	log.Fatal(http.ListenAndServe(":8080", JSONMiddleware(m)))
}
