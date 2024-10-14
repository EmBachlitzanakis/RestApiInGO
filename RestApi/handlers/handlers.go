package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"RestApi/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetBooks handler to fetch all books
func GetBooks(appState *model.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appState.Books.Lock()
		defer appState.Books.Unlock()

		// Encode and return the list of books
		err := json.NewEncoder(w).Encode(appState.BooksList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error encoding books: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// GetBook is an HTTP handler that fetches a book by ID.
func GetBook(appState *model.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid book ID: %v", err)
			return
		}

		appState.Books.Lock()
		defer appState.Books.Unlock()

		var book *model.Book
		for _, b := range appState.BooksList {
			if b.ID == id {
				book = &b
				break
			}
		}
		if book == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Book not found")
			return
		}

		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error encoding book: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// CreateBook creates a new book (POST /books)
func CreateBook(appState *model.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook model.Book
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		appState.Books.Lock()
		defer appState.Books.Unlock()

		newBook.ID = uuid.New()
		appState.BooksList = append(appState.BooksList, newBook)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBook)
	}
}

// UpdateBook updates a book by ID (PUT /books/{id})
func UpdateBook(appState *model.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID, err := uuid.Parse(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid book ID"))
			return
		}

		var updatedBook model.Book
		err = json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		appState.Books.Lock()
		defer appState.Books.Unlock()

		for i, book := range appState.BooksList {
			if book.ID == bookID {
				appState.BooksList[i].Title = updatedBook.Title
				appState.BooksList[i].Author = updatedBook.Author

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(appState.BooksList[i])
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
	}
}

// DeleteBook deletes a book by ID (DELETE /books/{id})
func DeleteBook(appState *model.AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID, err := uuid.Parse(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid book ID"))
			return
		}

		appState.Books.Lock()
		defer appState.Books.Unlock()

		for i, book := range appState.BooksList {
			if book.ID == bookID {
				appState.BooksList = append(appState.BooksList[:i], appState.BooksList[i+1:]...)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Book deleted"))
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
	}
}
