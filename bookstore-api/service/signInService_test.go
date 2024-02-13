package service

import (
	"bookstore-api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignInService_SignIn_Success(t *testing.T) {
	signInService := NewSignInService()

	validRequest := model.SignInRequest{
		Username: "testuser",
		Password: "testuser",
	}

	response, err := signInService.SignIn(validRequest)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Token)
}

func TestSignInService_SignIn_InvalidCredentials(t *testing.T) {
	signInService := NewSignInService()

	invalidRequest := model.SignInRequest{
		Username: "testuser",
		Password: "wrongpassword",
	}

	response, err := signInService.SignIn(invalidRequest)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.EqualError(t, err, "Invalid credentials")
}

func TestSignInService_GenerateToken_Success(t *testing.T) {
	signInService := NewSignInService()

	validRequest := model.SignInRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	token, err := signInService.generateToken(validRequest)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestSignInService_GenerateToken_Error(t *testing.T) {
	signInService := NewSignInService()

	invalidRequest := model.SignInRequest{
		Username: "testuser",
		Password: "",
	}

	token, err := signInService.generateToken(invalidRequest)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.EqualError(t, err, "Invalid request for token generation")
}
