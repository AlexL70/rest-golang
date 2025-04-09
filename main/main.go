package main

import "fmt"

func main() {
	// Initialize the database connection
	connStr := "host=localhost user=postgres password=postgres dbname=rest-golang port=5432 sslmode=disable"
	db, err := getDB(connStr)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %v", err))
	}

	addr := ":3000"
	srv := NewAPIServer(addr, db)
	srv.Run()
}
