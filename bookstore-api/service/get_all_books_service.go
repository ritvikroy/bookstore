package service

import (
	"bookstore-api/model"
	"bookstore-api/repository"
	"context"
	"errors"
)

//go:generate mockgen -source=get_all_books_service.go -destination=../mocks/get_all_books_service_mock.go -package=mocks
type bookstoreService struct {
	repository repository.BookStoreRepository
}

type BookStoreService interface {
	GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (model.AllBooks, error)
}

func NewBooksService(repo repository.BookStoreRepository) BookStoreService {
	return &bookstoreService{
		repository: repo,
	}
}

func (b *bookstoreService) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (model.AllBooks, error) {
	books := make([]model.Books, 0)
	books, err := b.repository.GetAllBooks(ctx, searchText, pageSize, pageNum)
	if err != nil {
		return model.AllBooks{}, errors.New("some error occures")
	}

	return model.AllBooks{Books: books}, nil
}
