package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Book struct {
	ID     primitive.ObjectID `bson:"_id" binding:"required"`
	Isbn   string             `bson:"isbn,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
