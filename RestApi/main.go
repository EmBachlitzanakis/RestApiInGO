package main

import (
	"log"
	"net/http"

	"RestApi/handlers"
	"RestApi/model"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the AppState with SQLite database
	appState, err := model.NewAppState("books.db")
	if err != nil {
		log.Fatalf("Failed to initialize app state: %v", err)
	}

	// Close the database connection when the program exits
	defer appState.DB.Close()

	// Set up the router
	r := mux.NewRouter()

	// Define your routes (CRUD operations)
	r.HandleFunc("/books", handlers.GetBooks(appState)).Methods("GET")           // Get all books
	r.HandleFunc("/books/{id}", handlers.GetBook(appState)).Methods("GET")       // Get a specific book by ID
	r.HandleFunc("/books", handlers.CreateBook(appState)).Methods("POST")        // Create a new book
	r.HandleFunc("/books/{id}", handlers.UpdateBook(appState)).Methods("PUT")    // Update a book by ID
	r.HandleFunc("/books/{id}", handlers.DeleteBook(appState)).Methods("DELETE") // Delete a book by ID

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
