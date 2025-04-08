package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	id_str := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(id_str, 10, 32)
	if err != nil {
		err_str := fmt.Sprintf("Invalid post ID: \"%s\". Error parsing the value: positive integer expected.", id_str)
		http.Error(w, err_str, http.StatusBadRequest)
		return
	}
	res, err := s.repo.getPost(id)
	if err != nil {
		err_str := fmt.Sprintf("Post with ID %d not found. Error: %v", id, err.Error())
		http.Error(w, err_str, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
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
