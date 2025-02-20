package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ajarodiy/jwt-auth-go/models"
	"github.com/ajarodiy/jwt-auth-go/utils"
)

var users = map[string]models.User{} // In-memory user store; replace with a database in production

func Register(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if _, exists := users[creds.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(creds.PasswordHash)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	users[creds.Username] = models.User{
		Username:     creds.Username,
		PasswordHash: hashedPassword,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, exists := users[creds.Username]
	if !exists || !utils.CheckPasswordHash(creds.PasswordHash, user.PasswordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true, // Set to true in production
	})

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
