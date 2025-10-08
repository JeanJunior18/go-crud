package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"

	model "github.com/JeanJunior18/go-crud/internal/domain"
	"github.com/JeanJunior18/go-crud/internal/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const postCollection = "posts"

type MongoPostRepository struct {
	client *mongo.Client
	dbName string
}

func New(client *mongo.Client, dbName string) *MongoPostRepository {
	log.Print("MONGO ADAPTER INITIALIZED")
	return &MongoPostRepository{
		client: client,
		dbName: dbName,
	}
}

var _ repository.PostRepository = (*MongoPostRepository)(nil)

func (r *MongoPostRepository) Save(post model.Post) error {
	ctx := context.TODO()

	collection := r.client.Database(r.dbName).Collection(postCollection)

	_, err := collection.InsertOne(ctx, post)

	if err != nil {
		return fmt.Errorf("error on insert post")
	}
	return nil
}

func (r *MongoPostRepository) FindByID(id uuid.UUID) (*model.Post, error) {
	return nil, errors.New("not implemented")
}

func (r *MongoPostRepository) FindAll() ([]model.Post, error) {
	return []model.Post{}, errors.New("not implemented")
}

func (r *MongoPostRepository) UpdatePost(id uuid.UUID, post model.Post) error {
	return errors.New("not implemented")
}

func (r *MongoPostRepository) DeletePost(id uuid.UUID) error {
	return errors.New("not implemented")
}
