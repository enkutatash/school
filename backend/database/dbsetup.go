package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client = DBsetup()
)

func DBsetup() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}


func GetData(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("highschool").Collection(collectionName)
}

