// 将之前的web服务的全局变量存储 var books = make(map[string]Book)
// 更换为数据库存储

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Stock  string
}

func main() {
	connStr := "postgres://postgres:123456@localhost:5432/muxi202507?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("连接成功！")

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
