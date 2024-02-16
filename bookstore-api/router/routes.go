package router

import (
	"bookstore-api/controllers"
	"bookstore-api/repository"
	"bookstore-api/service"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(db *sql.DB, appEngine *gin.Engine) *gin.Engine {
	signInService := service.NewSignInService()
	signInController := controllers.NewSignInController(signInService)
	bookRepository := repository.NewBooksRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	booksService := service.NewBooksService(bookRepository,orderRepo)
	booksStoreController := controllers.NewBooksController(booksService)

	appEngine.POST("/api/signin", signInController.HandleSignIn)
	apiRouter := appEngine.Group("/api")
	{
		authRouter := apiRouter.Group("/auth")
		{
			apiRouter.Use(middlewareFunc)
			authRouter.GET("/books", booksStoreController.GetAllBooks)
			authRouter.POST("/order/book", booksStoreController.BuyBook)
		}
	}

	return appEngine
}

func middlewareFunc(c *gin.Context) {
	values := c.Request.Header["Denvar"]
	sDec, _ := base64.StdEncoding.DecodeString(values[0])
	fmt.Println("userId: " ,strings.SplitAfter(string(sDec), "--")[2])

	c.Set("userId",strings.SplitAfter(string(sDec), "--")[2])
	c.Next()
}
