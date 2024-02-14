package repository

import (
	"bookstore-api/model"
	"context"
	"database/sql"
	"fmt"
)

type bookStoreRepository struct {
	dbConnection *sql.DB
}
type BookStoreRepository interface {
	GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (error, []model.Books)
}

func NewBooksRepository(dbConnection *sql.DB) BookStoreRepository {
	return bookStoreRepository{
		dbConnection: dbConnection,
	}
}

const (
	GetAllBooks = `select * from books`
)

func (b bookStoreRepository) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (error, []model.Books) {
	fmt.Println(" ----- bookStoreRepository")
	var books = make([]model.Books, 0)
	rows, err := b.dbConnection.Query(GetAllBooks)
	if err != nil {
		fmt.Println("Error in getting the row")
	}

	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		book := model.Books{}
		fmt.Println(" ----- ", rows)
		scanError := rows.Scan(&book.Id, &book.Name, &book.Price, &book.Description)
		if scanError != nil {
			fmt.Println("Error in getting the scanError", scanError)
		}
		books = append(books, book)
	}
	return nil, books
}
