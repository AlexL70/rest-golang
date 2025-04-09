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
