package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet).Methods("GET")

	fmt.Printf("Starting server on address %s\n", s.addr)
	http.ListenAndServe(s.addr, router)
}
