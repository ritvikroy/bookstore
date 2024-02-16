package controllers

import (
	"bookstore-api/model"
	"bookstore-api/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type booksController struct {
	service service.BookStoreService
}

type BooksController interface {
	GetAllBooks(ctx *gin.Context)
	BuyBook(ctx *gin.Context)
}

func NewGetAllBooks(service service.BookStoreService) BooksController {
	return booksController{
		service: service,
	}
}

func (b booksController) GetAllBooks(ctx *gin.Context) {
	var pageNumInt int
	var err error
	searchQuery, _ := ctx.GetQuery("searchText")
	limit, _ := ctx.GetQuery("limit")
	pageNum, _ := ctx.GetQuery("pageNum")
	fmt.Println(searchQuery, limit, pageNum)

	limitInt, limitIntErr := checkingPazeSizeAndPageNumber(limit)
	if limitIntErr != nil {
		fmt.Println("error converting  pageSizeInt string to int")
		return
	}
	// if pageNum != "" {
	// 	pageNumInt, err = strconv.Atoi(pageNum)
	// 	if err != nil {
	// 		fmt.Println("error converting pageNum string to int")
	// 		return
	// 	}
	// }
	pageNumInt, pageNumIntErr := checkingPazeSizeAndPageNumber(pageNum)
	if pageNumIntErr != nil {
		fmt.Println("error converting  pageSizeInt string to int")
		return
	}
	fmt.Println("Method: Controller : GetAllBooks")
	allBooks, err := b.service.GetAllBooks(ctx, searchQuery, limitInt, pageNumInt)
	if err != nil {
		fmt.Printf("error occured : %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, allBooks)
}

func (b booksController) BuyBook(ctx *gin.Context) {
	fmt.Println("Method:BuyBook, Class:booksController")
	var buyBookRequest model.BuyBookRequest
	if err := ctx.ShouldBindJSON(&buyBookRequest); err != nil {
		fmt.Println("error occured")
		ctx.JSON(http.StatusBadRequest, errors.New("invalid buy book input data"))
		return
	}

	orderId, err := b.service.BuyBook(ctx, buyBookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("order failed"))
		return
	}

	ctx.JSON(http.StatusOK, model.OrderResponse{
		OrderId: orderId,
	})
}

func checkingPazeSizeAndPageNumber(parameter string) (int, error) {
	if parameter != "" {
		parameterInt, err := strconv.Atoi(parameter)
		if err != nil {
			fmt.Println("error converting  pageSizeInt string to int", err)
			return 0, errors.New("error in parsing the parameter")
		}
		return parameterInt, nil
	}
	return 0, nil
}
