package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应体，用来转化gin.H给swag
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Book 书籍结构体
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

var books = make(map[string]Book) // 存储已有的book的map

// AddBook 添加书籍
// @Summary 添加书籍
// @Description 传入书籍信息新增书籍
// @Tags 图书
// @Accept json
// @Produce json
// @Param book body Book true "书籍信息"
// @Success 200 {object} Response{data=Book} "Book added"
// @Failure 400 {object} Response "400错误"
// @Failure 200 {object} Response "book already exists"
// @Router /books [post]
func AddBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
		return
	}

	if _, exists := books[book.ID]; exists {
		c.JSON(http.StatusOK, Response{Message: "book already exists"})
		return
	}

	books[book.ID] = book
	c.JSON(http.StatusOK, Response{Message: "Book added", Data: book})
}

// UpdateBook 更新书籍
// @Summary 更新书籍
// @Description 更新书籍
// @Tags 图书
// @Accept json
// @Produce json
// @Param id path string true "书籍ID"
// @Success 200 {object} Response{data=Book} "Book updated"
// @Failure 404 {object} Response "Book not found"
// @Failure 400 {object} Response "400错误"
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
		return
	}
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, Response{Message: "Book not found"})
		return
	}
	book.ID = id
	books[id] = book
	c.JSON(http.StatusOK, Response{Message: "Book updated", Data: book})
}

// DeleteBook 删除书籍
// @Summary 删除书籍
// @Description 删除书籍
// @Tags 图书
// @Accept json
// @Produce json
// @Param id path string true "书籍ID"
// @Success 200 {object} Response "Book deleted"
// @Failure 404 {object} Response "Book not exists"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, Response{Message: "book not exists"})
		return
	}
	delete(books, id)
	c.JSON(http.StatusOK, Response{Message: "book deleted"})
}

// SearchBook 搜索书籍
// @Summary 搜索书籍
// @Description 搜索书籍
// @Tags 图书
// @Accept json
// @Produce json
// @Param book body Book true "书籍信息"
// @Success 200 {object} Response{data=Book}
// @Failure 200 {object} Response{data=Book} "success"
// @Failure 404 {object} Response "book not found"
// @Router /books [get]
func SearchBook(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		booklist := make([]Book, 0, len(books))
		for _, book := range books {
			booklist = append(booklist, book)
		}
		c.JSON(http.StatusOK, Response{Message: "success", Data: booklist})
		return
	}
	if book, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, Response{Message: "book not found"})
		return
	} else {
		c.JSON(http.StatusOK, book)
	}
}

// @title Book Management API
// @version 1.0
// @description This is a sample server for managing books.
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()
	r.POST("/books", AddBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
	r.GET("/books", SearchBook)
	r.Run()
}
