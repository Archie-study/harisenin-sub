package models

import (
	"database/sql"
	"fmt"

	"github.com/Archie-study/harisenin-sub/config"
	"github.com/Archie-study/harisenin-sub/entities"
)

// Membuat struct BookModel
type BookModel struct {
	conn *sql.DB
}

// Untuk mengembalikan struct Book dalam bentuk pointer
func NewBookModel() *BookModel {
	conn, err := config.DBConnection()

	if err != nil {
		panic(err)
	}

	return &BookModel{
		conn: conn,
	}
}

// Struct method untuk mengambil semua data book
func (b *BookModel) FindAll() ([]entities.Book, error) {

	rows, err := b.conn.Query("select * from book")
	if err != nil {
		return []entities.Book{}, err
	}
	defer rows.Close()

	var dataBook []entities.Book
	for rows.Next() {
		var book entities.Book
		rows.Scan(&book.Id,
			&book.BookTitle,
			&book.BookAuthor,
			&book.BookYear,
			&book.BookDescription,
			&book.BookRating)

		if book.BookRating == "5" {
			book.BookRating = "⭐⭐⭐⭐⭐"
		} else if book.BookRating == "4" {
			book.BookRating = "⭐⭐⭐⭐"
		} else if book.BookRating == "3" {
			book.BookRating = "⭐⭐⭐"
		} else if book.BookRating == "2" {
			book.BookRating = "⭐⭐"
		} else {
			book.BookRating = "⭐"
		}

		dataBook = append(dataBook, book)
	}

	return dataBook, nil
}

// Struct untuk proses menyimpan data ke database
func (b *BookModel) Create(book entities.Book) bool {

	result, err := b.conn.Exec("insert into book (book_title, book_author, book_year, book_description, book_rating) values(?,?,?,?,?)",
		book.BookTitle, book.BookAuthor, book.BookYear, book.BookDescription, book.BookRating)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (b *BookModel) Find(id int64, book *entities.Book) error {

	return b.conn.QueryRow("select * from book where id = ?", id).Scan(
		&book.Id,
		&book.BookTitle,
		&book.BookAuthor,
		&book.BookYear,
		&book.BookDescription,
		&book.BookRating)
}

func (b *BookModel) Update(book entities.Book) error {

	_, err := b.conn.Exec(
		"update book set book_title = ?, book_author = ?, book_year = ?, book_description = ?, book_rating = ? where id = ?",
		book.BookTitle, book.BookAuthor, book.BookYear, book.BookDescription, book.BookRating, book.Id)

	if err != nil {
		return err
	}

	return nil
}

func (b *BookModel) Delete(id int64) {
	b.conn.Exec("delete from book where id = ?", id)
}
