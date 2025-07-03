# Go Book App

This project is a simple Go application that interacts with a PostgreSQL database to manage a collection of books. It allows users to add, update, delete, and search for books in the database.

## Project Structure

```
go-book-app
├── Dockerfile
├── docker-compose.yml
├── t1.go
└── README.md
```

## Requirements

- Go (version 1.16 or higher)
- Docker
- Docker Compose
- PostgreSQL

## Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd go-book-app
   ```

2. Build the Docker image:

   ```
   docker build -t go-book-app .
   ```

3. Start the application using Docker Compose:

   ```
   docker-compose up
   ```

## Usage

- The application exposes an API to manage books. You can use tools like Postman or curl to interact with the API.
- The following endpoints are available:
  - `POST /books` - Add a new book
  - `PUT /books/{id}` - Update an existing book
  - `DELETE /books/{id}` - Delete a book
  - `GET /books` - Retrieve all books or a specific book by ID

## Database Configuration

The application connects to a PostgreSQL database. Ensure that the database is running and the connection details are correctly specified in the `docker-compose.yml` file.

## License

This project is licensed under the MIT License.