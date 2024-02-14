package router

import (
	"bookstore-api/controllers"
	"bookstore-api/repository"
	"bookstore-api/service"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(db *sql.DB, appEngine *gin.Engine) *gin.Engine {
	signInService := service.NewSignInService()
	signInController := controllers.NewSignInController(signInService)
	bookRepository := repository.NewBooksRepository(db)
	booksService := service.NewBooksService(bookRepository)
	booksStoreController := controllers.NewGetAllBooks(booksService)

	appEngine.POST("/api/signin", signInController.HandleSignIn)
	appEngine.GET("/api/books", booksStoreController.GetAllBooks)
	return appEngine
}

func middlewareFunc(c *gin.Context) {
	fmt.Println("middleware")

	c.Next()
}
