package model

import (
	"sync"

	"github.com/google/uuid"
)

type Book struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

type AppState struct {
	Books     sync.Mutex
	BooksList []Book
}
