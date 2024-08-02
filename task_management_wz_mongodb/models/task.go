package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID 		primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     primitive.DateTime `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
}
