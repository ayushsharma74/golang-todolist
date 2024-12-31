package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	clientOps := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(context.Background(), clientOps)

	if err != nil {
		log.Fatal(err)
	}

	return client
}