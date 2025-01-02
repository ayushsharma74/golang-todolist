package task

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"todo-list-app/internal/task"
)

// RegisterRoutes sets up the routes for task management
func RegisterRoutes(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("/tasks", CreateTask(client)).Methods("POST")
	r.HandleFunc("/tasks", GetTasks(client)).Methods("GET")
	r.HandleFunc("/tasks/{id}", UpdateTask(client)).Methods("PUT")
	r.HandleFunc("/tasks/{id}", DeleteTask(client)).Methods("DELETE")
}

// CreateTask handles task creation
func CreateTask(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task task.Task
		// Decode request body into task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Call service to create the task
		service := task.NewService(client)
		err := service.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
