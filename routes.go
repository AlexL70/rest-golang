package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := &GreetRes{
		Hello: "Hello, World!",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleGetPost(w http.ResponseWriter, r *http.Request) {
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

func (s *APIServer) handleSavePost(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	// Read request body
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Copy payload to post
	post := &Post{
		Title:     payload.Title,
		Content:   payload.Content,
		CreatedAt: time.Now(),
	}

	// Save post to the database
	newPost, err := s.repo.createPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}
