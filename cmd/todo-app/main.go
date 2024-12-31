package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-list-app/pkg/mongo"

	"github.com/gorilla/mux"
)

func main() {
	client := mongo.ConnectDB()
	defer client.Disconnect()

	r := mux.NewRouter()

	fmt.Println("Server Started At PORT 8000")
	log.Fatal(http.ListenAndServe(":8080", r))

}