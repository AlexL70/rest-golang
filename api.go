package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	repo PostRepo
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		repo: &postRepo{
			db: db,
		},
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet).Methods("GET")
	router.HandleFunc("/post", s.handleGetPost).Methods("GET")
	router.HandleFunc("/post", s.handleSavePost).Methods("POST")

	fmt.Printf("Starting server on address %s\n", s.addr)
	http.ListenAndServe(s.addr, router)
}
