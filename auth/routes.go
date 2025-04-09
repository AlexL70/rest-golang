package main

import (
	"encoding/json"
	"net/http"
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
	res := &KLoginRes{
		AccessToken: "dummy_access_token",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
