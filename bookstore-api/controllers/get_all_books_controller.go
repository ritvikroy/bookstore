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
	var buyBookRequest model.BuyBookRequest
	if err := ctx.ShouldBindJSON(&buyBookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buy book input data"})
		return
	}

	err := b.service.BuyBook(buyBookRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	ctx.JSON(http.StatusOK)
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
