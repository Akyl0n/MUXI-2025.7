// @title           Book API
// @version         1.0
// @description     用于管理书籍的简单 RESTful API 示例
// @contact.name    muxi
// @contact.email   muxi@example.com
// @host      localhost:8080
// @BasePath  /

package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// Book 书籍结构体
// @Description 书籍信息
// @Param id string 书籍ID
// @Param title string 书名
// @Param author string 作者
// @Param stock string 库存

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

var db *sql.DB

func AddBook(db *sql.DB, book Book) error {
	_, err := db.Exec(`INSERT INTO books (id, title, author, stock) VALUES ($1, $2, $3, $4)`,
		book.ID, book.Title, book.Author, book.Stock)
	return err
}

func UpdateBook(db *sql.DB, book Book) error {
	_, err := db.Exec(`UPDATE books SET title=$1, author=$2, stock=$3 WHERE id=$4`,
		book.Title, book.Author, book.Stock, book.ID)
	return err
}

func DeleteBook(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM books WHERE id=$1`, id)
	return err
}

func SearchBook(db *sql.DB, id string) ([]Book, error) {
	var rows *sql.Rows
	var err error
	if id == "" {
		rows, err = db.Query(`SELECT id, title, author, stock FROM books`)
	} else {
		rows, err = db.Query(`SELECT id, title, author, stock FROM books WHERE id=$1`, id)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Stock); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// @Summary      获取所有书籍
// @Description  获取书籍列表
// @Tags         books
// @Produce      json
// @Success      200  {array}   Book
// @Failure      500  {object}  map[string]string
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	books, err := SearchBook(db, "")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

// @Summary      新增书籍
// @Description  添加一本新书
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      Book  true  "书籍信息"
// @Success      201   {string}  string  "ok"
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /books [post]
func AddBookHandler(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	if err := AddBook(db, book); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.String(201, "ok")
}

// @Summary      获取单本书籍
// @Description  根据ID获取书籍
// @Tags         books
// @Produce      json
// @Param        id   path      string  true  "书籍ID"
// @Success      200  {object}  Book
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	books, err := SearchBook(db, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(books) == 0 {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, books[0])
}

// @Summary      更新书籍
// @Description  根据ID更新书籍
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "书籍ID"
// @Param        book  body      Book    true  "书籍信息"
// @Success      200   {string}  string  "ok"
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /books/{id} [put]
func UpdateBookHandler(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	book.ID = id
	if err := UpdateBook(db, book); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.String(200, "ok")
}

// @Summary      删除书籍
// @Description  根据ID删除书籍
// @Tags         books
// @Produce      json
// @Param        id   path      string  true  "书籍ID"
// @Success      200  {string}  string  "ok"
// @Failure      500  {object}  map[string]string
// @Router       /books/{id} [delete]
func DeleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteBook(db, id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.String(200, "ok")
}

func main() {
	connStr := "postgres://postgres:123456@localhost:5432/muxi202507?sslmode=disable"
	var err error

	// 启动时重试数据库连接，最多重试10次
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		fmt.Println("等待数据库启动...", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("连接成功！")
	log.Println("[INFO] 数据库连接成功，准备建表...")

	// 自动建表
	tableSQL := `CREATE TABLE IF NOT EXISTS books (
		id VARCHAR(64) PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		stock VARCHAR(32) NOT NULL
	)`
	_, err = db.Exec(tableSQL)
	if err != nil {
		log.Fatalf("[ERROR] 建表失败: %s", err.Error())
	}
	log.Println("[INFO] 数据表 books 检查/创建完成")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	r.GET("/books", GetBooks)
	r.POST("/books", AddBookHandler)
	r.GET("/books/:id", GetBookByID)
	r.PUT("/books/:id", UpdateBookHandler)
	r.DELETE("/books/:id", DeleteBookHandler)

	log.Println("[INFO] 后端服务启动，监听端口 :8080")
	r.Run(":8080")
}
