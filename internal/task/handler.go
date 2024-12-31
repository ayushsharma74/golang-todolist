package task

import (
	"net/http"
	"todo-list-app/internal/task"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)
// TODO
func RegisterRoutes(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("/tasks", CreateTask(client)).Methods("POST")
}

func CreateTask(client *mongo.Client) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var task task.
	}
}