package main

import "net/http"

func handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
