package router

import (
	"bookstore-api/controllers"
	"bookstore-api/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes() *gin.Engine {
	signInService := service.NewSignInService()
	appEngine := gin.Default()
	authRoutes := appEngine.Group("v1")
	{
		authRoutes.Use(middlewareFunc)
		signInController := controllers.NewSignInController(signInService)

		appEngine.POST("/api/signin", signInController.HandleSignIn)

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
