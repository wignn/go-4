package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string      `json:"title"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
}
