package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

var collection *mongo.Collection
var ctx = context.TODO()

func OpenDatabase() error {
	log.Printf("Getting Db connection")
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_URL")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Connected to Database!")

	collection = client.Database("stores").Collection("stores")
	return nil
}

func GetCollection() *mongo.Collection {
	return collection
}
