package main

func main() {
	addr := ":4000"
	srv := NewAPIServer(addr)
	srv.Run()
}
