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
	userService := service.NewUserService()
	signInService := service.NewSignInService()
	fmt.Println("request -----")
	// appEngine := gin.Default()
	// authRoutes := appEngine.Group("/v1")
	//{

	userController := controllers.NewUser(userService)
	signInController := controllers.NewSignInController(signInService)
	bookRepository := repository.NewBooksRepository(db)
	booksService := service.NewBooksService(bookRepository)
	booksStoreController := controllers.NewGetAllBooks(booksService)

	appEngine.GET("/:id", userController.GetUserById)
	appEngine.POST("/api/signin", signInController.HandleSignIn)
	appEngine.GET("/api/books", booksStoreController.GetAllBooks)

	//}

	profileRoute := appEngine.Group("profile")
	{
		profileRoute.POST("/profile", controllers.UpdateProfile)
	}
	return appEngine
}

func middlewareFunc(c *gin.Context) {
	fmt.Println("middleware")

	c.Next()
}
