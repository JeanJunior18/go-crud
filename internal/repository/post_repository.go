package repository

import (
	"context"

	model "github.com/JeanJunior18/go-crud/internal/domain"
)

type PostRepository interface {
	Save(ctx context.Context, post model.Post) error
	FindByID(ctx context.Context, id string) (*model.Post, error)
	FindAll(ctx context.Context) ([]model.Post, error)
	UpdatePost(ctx context.Context, id string, post model.Post) error
	DeletePost(ctx context.Context, id string) error
}
