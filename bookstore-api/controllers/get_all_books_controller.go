package controllers

import (
	"bookstore-api/service"
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
}

func NewGetAllBooks(service service.BookStoreService) BooksController {
	return booksController{
		service: service,
	}
}

func (b booksController) GetAllBooks(ctx *gin.Context) {
	var pageNumInt, pageSizeInt int
	var err error
	searchQuery, _ := ctx.GetQuery("searchText")
	pageSize, _ := ctx.GetQuery("page_size")
	pageNum, _ := ctx.GetQuery("page_num")
	fmt.Println(pageNum, pageSize)
	if pageSize != "" {
		pageSizeInt, err = strconv.Atoi(pageSize)
		if err != nil {
			fmt.Println("error converting  pageSizeInt string to int")
			return
		}
	}
	if pageNum != "" {
		pageNumInt, err = strconv.Atoi(pageNum)
		if err != nil {
			fmt.Println("error converting pageNum string to int")
			return
		}
	}
	fmt.Println("Method: Controller : GetAllBooks")
	err, allBooks := b.service.GetAllBooks(ctx, searchQuery, pageSizeInt, pageNumInt)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	ctx.JSON(http.StatusOK, allBooks)
}
