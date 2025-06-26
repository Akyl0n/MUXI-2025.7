package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

var books = make(map[string]Book) // 存储已有的book的map

func AddBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := books[book.ID]; exists {
		c.JSON(http.StatusOK, gin.H{"error": "book already exists"})
		return
	}

	books[book.ID] = book
	c.JSON(http.StatusOK, gin.H{"message": "Book added", "book": book})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	book.ID = id
	books[id] = book
	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": book})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not exists"})
		return
	}
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func SearchBook(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		booklist := make([]Book, 0, len(books))
		for _, book := range books {
			booklist = append(booklist, book)
		}
		c.JSON(http.StatusOK, booklist)
	}
	if book, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func main() {
	r := gin.Default()
	r.POST("/books", AddBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
	r.GET("/books", SearchBook)
	r.Run()
}
