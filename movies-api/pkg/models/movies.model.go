package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID       primitive.ObjectID `json:"_id"`
	Title    string             `json:"title"`
	Ratings  float32            `json:"ratings"`
	Cast     []*Actor           `json:"cast"`
	Director *Director          `json:"director"`
}
