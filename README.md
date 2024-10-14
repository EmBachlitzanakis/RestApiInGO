# Simple Go REST API

This project implements a simple RESTful API for managing a collection of books using Go. It is designed for educational purposes to help you understand the basics of building a REST API with Go, utilizing concurrency, and managing shared state.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [API Endpoints](#api-endpoints)

## Overview

This REST API allows users to perform basic CRUD (Create, Read, Update, Delete) operations on a collection of books. It demonstrates how to handle HTTP requests concurrently and manage shared resources safely in Go.

## Features

- **Create a Book**: Add a new book to the collection.
- **Retrieve All Books**: Get a list of all books.
- **Retrieve a Book by ID**: Fetch details of a specific book using its ID.
- **Update a Book**: Modify the details of an existing book.
- **Delete a Book**: Remove a book from the collection.

## Technologies Used

- **Go**: The programming language used for building the API.
- **Gorilla Mux**: A powerful URL router and dispatcher for Golang.
- **UUID**: Used for generating unique identifiers for each book.
- **JSON**: Data format for sending and receiving data in the API.

## API Endpoints

| Method | Endpoint          | Description                     |
|--------|-------------------|---------------------------------|
| GET    | /books            | Retrieve all books              |
| GET    | /books/{id}       | Retrieve a specific book by ID  |
| POST   | /books            | Create a new book               |
| PUT    | /books/{id}       | Update an existing book         |
| DELETE | /books/{id}       | Delete a book by ID            |

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
