package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-secret-key")

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Token   string `json:"token"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var userRequest UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if userRequest.Name == "melika" && userRequest.Password == "test123" {

		claims := jwt.MapClaims{
			"username": userRequest.Name,
			"exp":      time.Now().Add(26 * time.Minute).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signedToken, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, " not generate", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Message: "successfull",
			Token:   signedToken,
		})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(Response{
		Error: "invalid",
	})
}