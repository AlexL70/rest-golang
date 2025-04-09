package main

import (
	"time"

	"gorm.io/gorm"
)

// GreetRes is the response structure for the greeting endpoint.
type GreetRes struct {
	Hello string `json:"hello"`
}

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Post struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"uniqueIndex"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
}

type PostRepo interface {
	createPost(post *Post) (*Post, error)
	getPost(id uint64) (*Post, error)
}

type postRepo struct {
	db *gorm.DB
}
