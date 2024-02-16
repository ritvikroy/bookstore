package repository

import (
	"bookstore-api/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type bookStoreRepository struct {
	dbConnection *sql.DB
}
type BookStoreRepository interface {
	GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) ([]model.Books, error)
	GetBookById(ctx context.Context, bookId string) (model.Books, error) 
}

func NewBooksRepository(dbConnection *sql.DB) BookStoreRepository {
	return &bookStoreRepository{
		dbConnection: dbConnection,
	}
}

const (
	GetAllBooks    = `select * from books`
	SearchBookText = `like '%' || $1 || '%'`
	limitBooks     = ` limit `
	offsetBooks    = ` offset `
	firstArgument  = `$1`
	secondArgument = `$2`
	thirdArgument  = `$3`

	GetBookById = `select * from books where id = $1`
	UpdateInventory = `update books set quantity= $1 where id= $2`
)

func (b *bookStoreRepository) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) ([]model.Books, error) {
	var rows *sql.Rows
	var err error
	var books = make([]model.Books, 0)
	if pageNum == 0 {
		pageNum = 0
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if searchText != "" {
		searchText = strings.ToLower(searchText)
		rows, err = b.dbConnection.Query(GetAllBooks+" where name "+SearchBookText+limitBooks+secondArgument+offsetBooks+thirdArgument, searchText, pageSize, pageNum)
	} else {

		rows, err = b.dbConnection.Query(GetAllBooks+limitBooks+firstArgument+offsetBooks+secondArgument, pageSize, pageNum)
	}
	if err != nil {
		fmt.Println("Error in getting the row", err)
	}

	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		book := model.Books{}
		
		scanError := rows.Scan(&book.Id, &book.Name, &book.Price, &book.Description, &book.Quantity)
		if scanError != nil {
			fmt.Println("Error in getting the scanError", scanError)
		}
		books = append(books, book)
	}
	return books, nil
}

func (b *bookStoreRepository) GetBookById(ctx context.Context, bookId string) (model.Books, error) {
	var book model.Books
	if err := b.dbConnection.QueryRow(GetBookById,
        bookId).Scan(&book); err != nil {
        if err == sql.ErrNoRows {
            return model.Books{}, errors.New("no rows for given bookId")
        }
        return model.Books{}, errors.New("some error occured")
    }
	return book, nil
}
