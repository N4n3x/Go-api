package main

import (
	"go-rest-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/books", models.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", models.GetBook).Methods("GET")
	r.HandleFunc("/api/books", models.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", models.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", models.DeleteBook).Methods("DELETE")

	// set our port address
	log.Fatal(http.ListenAndServe(":3000", r))

}
