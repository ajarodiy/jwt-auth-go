package main

import (
	"log"
	"net/http"

	"github.com/ajarodiy/jwt-auth-go/handlers"
	"github.com/ajarodiy/jwt-auth-go/middleware"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	protected := http.NewServeMux()
	protected.HandleFunc("/protected", ProtectedEndpoint)
	http.Handle("/protected", middleware.AuthMiddleware(protected))

	log.Println("Server starting on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "This is a protected endpoint"}`))
}
