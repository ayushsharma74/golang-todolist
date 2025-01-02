package task

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository interacts with MongoDB
type Repository struct {
	client *mongo.Client
}

// NewRepository creates a new task repository
func NewRepository(client *mongo.Client) *Repository {
	return &Repository{client: client}
}

// Create a new task in MongoDB
func (r *Repository) Create(task Task) error {
	collection := r.client.Database("todo_db").Collection("tasks")
	_, err := collection.InsertOne(context.Background(), task)
	return err
}
