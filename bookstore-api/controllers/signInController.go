package controllers

import (
	"bookstore-api/model"
	"bookstore-api/service"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInController struct {
	signInService *service.SignInService
}

func NewSignInController(signInService *service.SignInService) *SignInController {
	
	
	return &SignInController{
		signInService: signInService,
	}
}

// HandleSignIn handles the sign-in request.
func (sc *SignInController) HandleSignIn(ctx *gin.Context) {
	var signInRequest model.SignInRequest
	if err := ctx.ShouldBindJSON(&signInRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sign-in input data"})
		return
	}

	// Call the service to perform authentication
	signInResponse, err := sc.signInService.SignIn(signInRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	ctx.JSON(http.StatusOK, signInResponse)
}
