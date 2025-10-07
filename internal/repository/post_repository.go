package repository

import (
	model "github.com/JeanJunior18/go-crud/internal/domain"
	"github.com/google/uuid"
)

type PostRepository interface {
	Save(post model.Post) error
	FindByID(id uuid.UUID) (*model.Post, error)
	FindAll() ([]model.Post, error)
	UpdatePost(id uuid.UUID, post model.Post) error
	DeletePost(id uuid.UUID) error
}
