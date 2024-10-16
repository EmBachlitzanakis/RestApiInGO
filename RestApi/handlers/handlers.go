package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"RestApi/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetBooks handler to fetch all books from the SQLite database
func GetBooks(server *model.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Query all books from the database
		rows, err := server.DB.QueryContext(ctx, "SELECT id, title, author FROM books")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error querying books: %v", err)
			return
		}
		defer rows.Close()

		var books []model.Book

		for rows.Next() {
			var book model.Book
			err := rows.Scan(&book.ID, &book.Title, &book.Author)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error scanning books: %v", err)
				return
			}
			books = append(books, book)
		}

		// Encode and return the list of books
		err = json.NewEncoder(w).Encode(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error encoding books: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// CreateBook handler to add a new book to the SQLite database
func CreateBook(server *model.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var book model.Book

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error decoding book: %v", err)
			return
		}

		book.ID = uuid.New().String()

		_, err = server.DB.ExecContext(ctx, "INSERT INTO books (id, title, author) VALUES (?, ?, ?)", book.ID, book.Title, book.Author)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error inserting book: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	}
}

// GetBook handler to fetch a single book by ID
func GetBook(server *model.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		bookID := vars["id"]

		row := server.DB.QueryRowContext(ctx, "SELECT id, title, author FROM books WHERE id = ?", bookID)

		var book model.Book
		err := row.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Book not found: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

// UpdateBook updates a book by ID (PUT /books/{id})
func UpdateBook(server *model.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		vars := mux.Vars(r)
		bookID := vars["id"]

		var updatedBook model.Book
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		updatedBook.ID = bookID

		_, err = server.DB.ExecContext(ctx,
			"UPDATE books SET title = ?, author = ? WHERE id = ?",
			updatedBook.Title, updatedBook.Author, bookID,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error updating book: %v", err)
			return
		}

		result, err := server.DB.Exec("SELECT changes()")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error checking update result"))
			return
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedBook)
	}
}

// DeleteBook deletes a book by ID (DELETE /books/{id})
func DeleteBook(server *model.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		vars := mux.Vars(r)
		bookID := vars["id"]

		result, err := server.DB.ExecContext(ctx, "DELETE FROM books WHERE id = ?", bookID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error deleting book: %v", err)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error checking delete result"))
			return
		}
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Book deleted"))
	}
}
