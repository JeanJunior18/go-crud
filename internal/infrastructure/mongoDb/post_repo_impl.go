package mongodb

import (
	"errors"

	model "github.com/JeanJunior18/go-crud/internal/domain"
	"github.com/JeanJunior18/go-crud/internal/repository"
	"github.com/google/uuid"
)

type MongoPostRepositoryStub struct {
}

func New() *MongoPostRepositoryStub {
	return &MongoPostRepositoryStub{}
}

var _ repository.PostRepository = (*MongoPostRepositoryStub)(nil)

func (r *MongoPostRepositoryStub) Save(post model.Post) error {
	return nil
}

func (r *MongoPostRepositoryStub) FindByID(id uuid.UUID) (*model.Post, error) {
	return nil, nil
}

func (r *MongoPostRepositoryStub) FindAll() ([]model.Post, error) {
	return []model.Post{}, nil
}

func (r *MongoPostRepositoryStub) UpdatePost(id uuid.UUID, post model.Post) error {
	return errors.New("update não implementado no stub")
}

func (r *MongoPostRepositoryStub) DeletePost(id uuid.UUID) error {
	return nil
}
