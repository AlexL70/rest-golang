package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := &GreetRes{
		Hello: "Hello, World!",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	payload := &LoginPayload{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		fmt.Printf("Failed to decode request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Get clientId from the environment variable
	clientSecret := os.Getenv("KEYCLOACK_CLIENT_SECRET_FOR_GOLANG_REST")
	resp, err := s.kcClient.login(&KLoginPayload{
		clientId:     "rest-golang-auth",
		userName:     payload.Username,
		password:     payload.Password,
		grantType:    "password",
		clientSecret: clientSecret,
	})
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		formattedError := fmt.Sprintf("Login failed: %v", err)
		http.Error(w, formattedError, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
