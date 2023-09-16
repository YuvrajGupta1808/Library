package models

import (
	"database/sql"
	"myproject/config"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books []Book

func CreateBook(book Book) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	defer db.Close() // Ensure that the database connection is closed when the function returns

	query := "INSERT INTO book (id, title, author, quantity) VALUES ($1,$2,$3,$4)"
	_, err = db.Exec(query, book.ID, book.Title, book.Author, book.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByID(id int) (*Book, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	query := "SELECT id, title, author, quantity FROM book WHERE id = $1"
	row := db.QueryRow(query, id)

	var book Book
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Book not found
		}
		return nil, err
	}

	return &book, nil
}

func UpdateBook(id string, book Book) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()
	query := "UPDATE book SET title = $1, author = $3, quantity = $4 WHERE id = $5"
	_, err = db.Exec(query, book.Title, book.Author, book.Quantity, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()
	query := "DELETE FROM book WHERE id = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllBooks() ([]Book, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, title, author, quantity FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
