package router

import (
	"bookstore-api/controllers"
	"bookstore-api/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes() *gin.Engine {
	userService := service.NewUserService()
	signInService := service.NewSignInService()
	appEngine := gin.Default()
	authRoutes := appEngine.Group("v1")
	{
		// authRoutes.Use(middlewareFunc)
		userController := controllers.NewUser(userService)
		signInController := controllers.NewSignInController(signInService)

		authRoutes.GET("/:id", userController.GetUserById)
		authRoutes.POST("/api/signin", signInController.HandleSignIn)

	}

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
