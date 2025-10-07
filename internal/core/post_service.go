package core

import (
	"errors"
	"time"

	model "github.com/JeanJunior18/go-crud/internal/domain"
	"github.com/JeanJunior18/go-crud/internal/handler/dto"
	"github.com/JeanJunior18/go-crud/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type PostServiceContract interface {
	CreatePost(ctx *context.Context, data dto.CreatePostRequest) (model.Post, error)
}

type PostService struct {
	repo repository.PostRepository
}

func New(r repository.PostRepository) *PostService {
	return &PostService{
		repo: r,
	}
}

func (s *PostService) CreatePost(ctx *context.Context, data dto.CreatePostRequest) (model.Post, error) {
	if data.Title == "" || data.Content == "" {
		return model.Post{}, errors.New("missing data")
	}

	newPost := model.Post{
		ID:        uuid.New(),
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: time.Now(),
	}

	err := s.repo.Save(newPost)

	if err != nil {
		return model.Post{}, errors.New("error on save")
	}

	return newPost, nil
}
