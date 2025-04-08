package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", handleGreet)
	router.HandleFunc("/post", handlePost)

	fmt.Printf("Starting server on address %s\n", s.addr)
	http.ListenAndServe(s.addr, router)
}
