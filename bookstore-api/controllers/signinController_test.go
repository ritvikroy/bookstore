package controllers

import (
	"bookstore-api/model"
	"bookstore-api/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleSignIn_Success(t *testing.T) {
	// Setup
	mockSignInService := &service.SignInService{}
	signInController := NewSignInController(mockSignInService)

	signInRequest := model.SignInRequest{
		Username: "testuser",
		Password: "testuser",
	}

	jsonBody, _ := json.Marshal(signInRequest)
	req, err := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	signInController.HandleSignIn(ctx)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.SignInResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.NotEmpty(t, response.Token)
}

func TestHandleSignIn_InvalidInput(t *testing.T) {
	mockSignInService := &service.SignInService{}
	signInController := NewSignInController(mockSignInService)

	invalidSignInRequest := model.SignInRequest{}
	jsonBody, _ := json.Marshal(invalidSignInRequest)
	req, err := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	signInController.HandleSignIn(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid sign-in input data")
}

func TestHandleSignIn_AuthenticationFailure(t *testing.T) {
	mockSignInService := &service.SignInService{}
	signInController := NewSignInController(mockSignInService)

	validSignInRequest := model.SignInRequest{
		Username: "testuser",
		Password: "incorrectpassword",
	}
	jsonBody, _ := json.Marshal(validSignInRequest)
	req, err := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	signInController.HandleSignIn(ctx)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authentication failed")
}
