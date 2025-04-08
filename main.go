package main

func main() {
	addr := ":3000"
	srv := NewAPIServer(addr)
	srv.Run()
}
