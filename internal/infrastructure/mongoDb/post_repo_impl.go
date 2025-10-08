package mongodb

import (
	"context"
	"errors"
	"fmt"

	model "github.com/JeanJunior18/go-crud/internal/domain"
	"github.com/JeanJunior18/go-crud/internal/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const postCollection = "posts"

type MongoPostRepository struct {
	client *mongo.Client
	dbName string
}

func New(client *mongo.Client, dbName string) *MongoPostRepository {
	return &MongoPostRepository{
		client: client,
		dbName: dbName,
	}
}

var _ repository.PostRepository = (*MongoPostRepository)(nil)

func (r *MongoPostRepository) Save(ctx context.Context, post model.Post) error {

	collection := r.client.Database(r.dbName).Collection(postCollection)

	_, err := collection.InsertOne(ctx, post)

	if err != nil {
		return fmt.Errorf("error on insert post")
	}
	return nil
}

func (r *MongoPostRepository) FindByID(ctx context.Context, id string) (*model.Post, error) {
	postUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format provided: %w", err)
	}

	collection := r.client.Database(r.dbName).Collection(postCollection)
	result := collection.FindOne(ctx, bson.D{{"id", postUUID}})
	post := &model.Post{}

	if err := result.Decode(post); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to decode post with ID %s: %w", id, err)
	}
	return post, nil
}

func (r *MongoPostRepository) FindAll(ctx context.Context) ([]model.Post, error) {
	collection := r.client.Database(r.dbName).Collection(postCollection)

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error calling Find")
	}

	defer cursor.Close(ctx)

	var posts []model.Post

	for cursor.Next(ctx) {
		var post model.Post

		if err := cursor.Decode(&post); err != nil {
			return nil, fmt.Errorf("error deconding document: %w", err)
		}

		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor interation error: %w", err)
	}

	return posts, nil
}

func (r *MongoPostRepository) UpdatePost(ctx context.Context, id string, post model.Post) error {
	return errors.New("not implemented")
}

func (r *MongoPostRepository) DeletePost(ctx context.Context, id string) error {
	return errors.New("not implemented")
}
