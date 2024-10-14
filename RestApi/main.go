package main

import (
	"log"
	"net/http"
	"sync"

	"RestApi/handlers" // Adjust the import path based on your project structure
	"RestApi/model"

	"github.com/gorilla/mux"
)

// Initialize and set up the HTTP server
func main() {
	// Create the initial AppState with an empty book list and a mutex
	appState := &model.AppState{
		Books:     sync.Mutex{},
		BooksList: []model.Book{}, // Start with an empty book list
	}

	// Initialize the router
	router := mux.NewRouter()

	// Set up the routes and link to handlers, passing AppState
	router.HandleFunc("/books", handlers.GetBooks(appState)).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook(appState)).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook(appState)).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBook(appState)).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook(appState)).Methods("DELETE")

	// Start the HTTP server
	log.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
