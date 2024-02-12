package controllers

import (
	"bookstore-api/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUser(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u UserController) GetUserById(ctx *gin.Context) {

	userId := ctx.Param("id")

	str, err := u.userService.GetUserById(userId)
	fmt.Println(str, err)
}
