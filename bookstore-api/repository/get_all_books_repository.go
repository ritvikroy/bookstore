package repository

import (
	"bookstore-api/model"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type bookStoreRepository struct {
	dbConnection *sql.DB
}
type BookStoreRepository interface {
	GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) ([]model.Books,error)
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
)

func (b *bookStoreRepository) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) ([]model.Books, error) {
	fmt.Println(" ----- bookStoreRepository", searchText)
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
		fmt.Println(" ----- ", rows)
		scanError := rows.Scan(&book.Id, &book.Name, &book.Price, &book.Description)
		if scanError != nil {
			fmt.Println("Error in getting the scanError", scanError)
		}
		books = append(books, book)
	}
	return books,nil
}
