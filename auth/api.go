package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr       string
	httpClient *http.Client
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:       addr,
		httpClient: &http.Client{},
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet).Methods("GET")
	router.HandleFunc("/login", s.handleLogin).Methods("POST")

	fmt.Printf("Starting server on address %s\n", s.addr)
	http.ListenAndServe(s.addr, router)
}
