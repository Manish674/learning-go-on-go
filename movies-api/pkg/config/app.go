package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MovieCollection *mongo.Collection
var ActorCollection *mongo.Collection
var DirectorCollection *mongo.Collection

func Connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("mongodb connection successful")

	MovieCollection = client.Database("moviesdb").Collection("movie")
	ActorCollection = client.Database("moviesdb").Collection("actor")
	DirectorCollection = client.Database("moviesdb").Collection("Director")
}
