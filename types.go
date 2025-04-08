package main

import "time"

// GreetRes is the response structure for the greeting endpoint.
type GreetRes struct {
	Hello string `json:"hello"`
}

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"uniqueIndex"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
}
