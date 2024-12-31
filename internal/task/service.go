package task

import (
	"go.mongodb.org/mongo-driver/mongo"
	"todo-list-app/internal/task"
)

// Service contains the repository methods
type Service struct {
	repository *task.Repository
}

// NewService creates a new task service
func NewService(client *mongo.Client) *Service {
	repository := task.NewRepository(client)
	return &Service{repository: repository}
}

// Create a new task
func (s *Service) Create(task task.Task) error {
	return s.repository.Create(task)
}
