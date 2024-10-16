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
	server, err := model.NewServer("books.db")
	if err != nil {
		log.Fatalf("Failed to initialize app state: %v", err)
	}

	// Close the database connection when the program exits
	defer server.DB.Close()

	// Set up the router
	r := mux.NewRouter()

	// Define your routes (CRUD operations)
	r.HandleFunc("/books", handlers.GetBooks(server)).Methods("GET")           // Get all books
	r.HandleFunc("/books/{id}", handlers.GetBook(server)).Methods("GET")       // Get a specific book by ID
	r.HandleFunc("/books", handlers.CreateBook(server)).Methods("POST")        // Create a new book
	r.HandleFunc("/books/{id}", handlers.UpdateBook(server)).Methods("PUT")    // Update a book by ID
	r.HandleFunc("/books/{id}", handlers.DeleteBook(server)).Methods("DELETE") // Delete a book by ID

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
