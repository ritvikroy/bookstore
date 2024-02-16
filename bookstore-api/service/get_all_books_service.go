package service

import (
	"bookstore-api/model"
	"bookstore-api/repository"
	"context"
	"errors"
	"fmt"
)

//go:generate mockgen -source=get_all_books_service.go -destination=../mocks/get_all_books_service_mock.go -package=mocks
type bookstoreService struct {
	repository repository.BookStoreRepository
	ordersRepo repository.OrdersRepository
}

type BookStoreService interface {
	GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (model.AllBooks, error)
	BuyBook(ctx context.Context, buyBookReq model.BuyBookRequest) (string, error)
}

func NewBooksService(repo repository.BookStoreRepository) BookStoreService {
	return &bookstoreService{
		repository: repo,
	}
}

func (b *bookstoreService) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) (model.AllBooks, error) {
	books, err := b.repository.GetAllBooks(ctx, searchText, pageSize, pageNum)
	if err != nil {
		return model.AllBooks{}, errors.New("some error occures")
	}

	return model.AllBooks{Books: books}, nil
}

func (b *bookstoreService) BuyBook(ctx context.Context, buyBookReq model.BuyBookRequest) (string, error) {
	fmt.Println("Method:BuyBook, Class:bookstoreService")
	bookDetails, err := b.repository.GetBookById(ctx, buyBookReq.BookId)
	if err != nil {
		return "", errors.New("some error occured")
	}
	//check book quantity exist
	if bookDetails.Quantity < buyBookReq.Quantity {
		return "", errors.New("invalid quantity, Please check available quantity")
	}

	updatedCount := bookDetails.Quantity - buyBookReq.Quantity
	orderValue := fmt.Sprintf("%v", bookDetails.Price*buyBookReq.Quantity)
	orderId, err := b.ordersRepo.CreateOrder(ctx, buyBookReq.BookId, orderValue, buyBookReq.Quantity, updatedCount)
	if err != nil {
		fmt.Println("errror :", err)
		return "", errors.New("unable to process you order")
	}


	fmt.Println("orderId : ", orderId)
	return orderId, nil
}
