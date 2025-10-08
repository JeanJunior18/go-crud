package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	model "github.com/JeanJunior18/go-crud/internal/domain"
	"github.com/JeanJunior18/go-crud/internal/handler/dto"
	"github.com/JeanJunior18/go-crud/internal/repository"
	"github.com/google/uuid"
)

type PostServiceContract interface {
	CreatePost(ctx context.Context, data dto.CreatePostRequest) (model.Post, error)
	FindPost(ctx context.Context) ([]model.Post, error)
	FindById(ctx context.Context, id string) (*model.Post, error)
}

type PostService struct {
	repo repository.PostRepository
}

func New(r repository.PostRepository) *PostService {
	return &PostService{
		repo: r,
	}
}

func (s *PostService) CreatePost(ctx context.Context, data dto.CreatePostRequest) (model.Post, error) {
	if data.Title == "" || data.Content == "" {
		return model.Post{}, errors.New("missing data")
	}

	newPost := model.Post{
		ID:        uuid.New(),
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: time.Now(),
	}

	err := s.repo.Save(ctx, newPost)

	if err != nil {
		return model.Post{}, fmt.Errorf("Error on save post: %w", err)
	}

	return newPost, nil
}

func (s *PostService) FindPost(ctx context.Context) ([]model.Post, error) {
	posts, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("Error on find one: %w", err)
	}

	return posts, nil
}

func (s *PostService) FindById(ctx context.Context, id string) (*model.Post, error) {
	post, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("Error on find id %s: %w", id, err)
	}

	if post == nil {
		return nil, nil
	}

	return post, nil
}
