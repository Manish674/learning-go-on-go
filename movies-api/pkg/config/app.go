package app

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
  if err != nil {
    log.Fatal(err)
  }

  fmt.Print("mongodb connection successful")
  collection = client.Database("moviesdb").Collection("movie")

  fmt.Print("Collection is ready")

}
