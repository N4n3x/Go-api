package models

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-api/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

// GetBooks
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var books []Book

	//Connection mongoDB with helper class
	collection := helper.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book Book
		// & character returns the memory address of the following variable.
		err := cur.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = json.NewEncoder(w).Encode(books) // encode similar to serialize process.
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var book Book
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	// connect db
	collection := helper.ConnectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book Book

	collection := helper.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	book.ID = id

	_ = json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	_ = json.NewEncoder(w).Encode(deleteResult)
}
