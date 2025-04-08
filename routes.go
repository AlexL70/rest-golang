package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := &GreetRes{
		Hello: "Hello, World!",
	}
	json.NewEncoder(w).Encode(res)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Handle the request to create a new post
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := &Post{
		ID:        1,
		Title:     "Sample Post",
		Content:   "This is a sample post content.",
		CreatedAt: time.Now(),
	}
	json.NewEncoder(w).Encode(res)
}
