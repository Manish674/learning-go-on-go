package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var movieCollection *mongo.Collection
var actorCollection *mongo.Collection
var directorCollection *mongo.Collection

func connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("mongodb connection successful")

	movieCollection = client.Database("moviesdb").Collection("movie")
	actorCollection = client.Database("moviesdb").Collection("actor")
	directorCollection = client.Database("moviesdb").Collection("Director")

	fmt.Print("Collection is ready")
}
