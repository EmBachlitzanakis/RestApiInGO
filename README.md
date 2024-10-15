# Simple Go REST API with SQLite

This project implements a simple RESTful API for managing a collection of books using Go and SQLite. It is designed for educational purposes to help you understand the basics of building a REST API with Go, integrating an SQLite database, and managing persistent data.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [UUID in the Project](#uuid-in-the-project)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Setup Instructions](#setup-instructions)
- [Conclusion](#conclusion)

## Overview

This REST API allows users to perform basic CRUD (Create, Read, Update, Delete) operations on a collection of books. Unlike the previous in-memory implementation, this version stores data persistently using an SQLite database. The API demonstrates how to handle HTTP requests, perform database operations such as querying and modifying data, and manage shared resources efficiently.

## Features

- **Create a Book**: Add a new book to the SQLite database.
- **Retrieve All Books**: Get a list of all books stored in the database.
- **Retrieve a Book by ID**: Fetch details of a specific book using its ID from the database.
- **Update a Book**: Modify the details of an existing book in the database.
- **Delete a Book**: Remove a book from the database by its ID.

## Technologies Used

- **Go**: The programming language used for building the API.
- **SQLite**: A lightweight, serverless database used to store book records persistently.
- **Gorilla Mux**: A powerful URL router and dispatcher for Golang.
- **UUID**: Used for generating unique identifiers for each book.
- **JSON**: Data format for sending and receiving data in the API.

## UUID in the Project

UUID (Universally Unique Identifier) is used in this project to assign a unique identifier to each book. This ensures that every book in the system can be uniquely identified, regardless of the title or author. 

We use the UUID library from `github.com/google/uuid`, which provides a straightforward way to generate UUIDs. Each time a new book is created, a version 4 (randomly generated) UUID is assigned to it. This guarantees that even in a distributed system, where multiple instances of the API may be running, the generated IDs will remain unique without any central coordination.

Example of a UUID:
```text
550e8400-e29b-41d4-a716-446655440000
```

## API Endpoints

| Method | Endpoint          | Description                     |
|--------|-------------------|---------------------------------|
| GET    | /books            | Retrieve all books              |
| GET    | /books/{id}       | Retrieve a specific book by ID  |
| POST   | /books            | Create a new book               |
| PUT    | /books/{id}       | Update an existing book         |
| DELETE | /books/{id}       | Delete a book by ID             |

### Request and Response Formats

- **Request**: JSON format for book creation and updates. Example:
    ```json
    {
        "title": "Book Title",
        "author": "Author Name"
    }
    ```

- **Response**: JSON format for all endpoints. Example:
    ```json
    {
        "id": "unique-book-id",
        "title": "Book Title",
        "author": "Author Name"
    }
    ```

## Database Schema

The SQLite database has a single table called `books` with the following schema:

```sql
CREATE TABLE IF NOT EXISTS books (
    id TEXT PRIMARY KEY,
    title TEXT,
    author TEXT
);
```

## Setup Instructions

### Prerequisites

1. **Go**: You need to have Go installed. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. **SQLite**: No external installation is required, as SQLite is embedded and works out-of-the-box with the Go package `github.com/mattn/go-sqlite3`.

### Steps to Run

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/simple-go-rest-api-sqlite.git
    cd simple-go-rest-api-sqlite
    ```

2. **Install the dependencies**:
    ```bash
    go mod tidy
    ```

3. **Run the server**:
    ```bash
    go run main.go
    ```

   The server will start on port `8080`.

4. **Database File**: The database will be automatically created as `books.db` in the project root when you run the application for the first time.

## Conclusion

This project demonstrates how to build a simple REST API in Go with persistent storage using SQLite. By using Gorilla Mux for routing and handling JSON requests/responses, the application provides a lightweight and extensible base for more complex RESTful services.
