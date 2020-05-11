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

type Character struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id" binding:"required"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Class      string             `json:"class,omitempty" bson:"class,omitempty"`
	Experience int64              `json:"experience,omitempty" bson:"experience,omitempty"`
	Stats      *Stats             `json:"stats,omitempty" bson:"stats,omitempty"`
}

type Stats struct {
	Strength     int `json:"strength,omitempty" bson:"strength,omitempty"`
	Dexterity    int `json:"dexterity,omitempty" bson:"dexterity,omitempty"`
	Intelligence int `json:"intelligence,omitempty" bson:"intelligence,omitempty"`
	Constitution int `json:"constitution,omitempty" bson:"constitution,omitempty"`
}

type Item struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}
