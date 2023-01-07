package models

import (
	"context"
	"movies-api/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Actor struct {
	ID        primitive.ObjectID `json:"_id"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Age       int                `json:"age"`
}

func GetActor(Id string) *mongo.SingleResult {
	app.Connect()
	var result *mongo.SingleResult
	app.ActorCollection.FindOne(context.Background(), bson.M{"firstname": Id}).Decode(&result)
	return result
}
