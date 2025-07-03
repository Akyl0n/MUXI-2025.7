package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
	"time"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

var db *sql.DB

func main() {
	connStr := "postgres://postgres:123456@db:5432/muxi202507?sslmode=disable"
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

	// 自动建表
	tableSQL := `CREATE TABLE IF NOT EXISTS books (
		id VARCHAR(64) PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		stock VARCHAR(32) NOT NULL
	)`
	_, err = db.Exec(tableSQL)
	if err != nil {
		panic("建表失败: " + err.Error())
	}

	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookHandler)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

// booksHandler 处理 /books 路由
func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 查询所有书籍
		books, err := SearchBook(db, "")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(books)
	case http.MethodPost:
		// 添加书籍
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}
		if err := AddBook(db, book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// bookHandler 处理 /books/{id} 路由
func bookHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing id"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		books, err := SearchBook(db, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if len(books) == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}
		json.NewEncoder(w).Encode(books[0])
	case http.MethodPut:
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}
		book.ID = id
		if err := UpdateBook(db, book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("ok"))
	case http.MethodDelete:
		if err := DeleteBook(db, id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("ok"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// AddBook 向数据库添加书籍
func AddBook(db *sql.DB, book Book) error {
	_, err := db.Exec(`INSERT INTO books (id, title, author, stock) VALUES ($1, $2, $3, $4)`,
		book.ID, book.Title, book.Author, book.Stock)
	return err
}

// UpdateBook 更新数据库中的书籍
func UpdateBook(db *sql.DB, book Book) error {
	_, err := db.Exec(`UPDATE books SET title=$1, author=$2, stock=$3 WHERE id=$4`,
		book.Title, book.Author, book.Stock, book.ID)
	return err
}

// DeleteBook 删除数据库中的书籍
func DeleteBook(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM books WHERE id=$1`, id)
	return err
}

// SearchBook 根据ID查找书籍，如果id为空则返回所有书籍
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
