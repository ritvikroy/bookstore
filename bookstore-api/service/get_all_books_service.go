package service

import (
	"bookstore-api/model"
	"bookstore-api/repository"
	"context"
	"errors"
)

type bookstoreService struct {
	repository repository.BookStoreRepository
}

type BookStoreService interface {
	GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (error, model.AllBooks)
}

func NewBooksService(repo repository.BookStoreRepository) BookStoreService {
	return bookstoreService{
		repository: repo,
	}
}

func (b bookstoreService) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (error, model.AllBooks) {
	books := make([]model.Books, 0)
	err, books := b.repository.GetAllBooks(ctx, searchText, pageSize, pageNum)
	if err != nil {
		return errors.New("some error occures"), model.AllBooks{}
	}

	return nil, model.AllBooks{Books: books}
}
