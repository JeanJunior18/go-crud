package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UserId    uuid.UUID `json:"userId"`
}

type PostService interface {
	CreatePost(post Post) error
	GetPost(id uuid.UUID) (*Post, error)
	ListPosts() ([]Post, error)
	UpdatePost(id uuid.UUID, post Post) error
	DeletePost(id uuid.UUID) error
}
