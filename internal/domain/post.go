package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id" bson:"id,type=string"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UserId    uuid.UUID `json:"userId" bson:"userId,type=string"`
}

type PostService interface {
	CreatePost(post Post) error
	GetPost(id uuid.UUID) (*Post, error)
	ListPosts() ([]Post, error)
	UpdatePost(id uuid.UUID, post Post) error
	DeletePost(id uuid.UUID) error
}
