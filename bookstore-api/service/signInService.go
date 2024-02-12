package service

import (
	"bookstore-api/model"
	"encoding/base64"
	"errors"
)

type SignInService struct {
}

func NewSignInService() *SignInService {
	return &SignInService{}
}

// SignIn performs user authentication.
func (ss *SignInService) SignIn(request model.SignInRequest) (*model.SignInResponse, error) {
	isValidCredentials := ss.authenticateUser(request.Username, request.Password)
	if !isValidCredentials {
		return nil, errors.New("Invalid credentials")
	}

	token, err := ss.generateToken(request)
	if err != nil {
		return nil, err
	}

	return &model.SignInResponse{Token: token}, nil
}

func (ss *SignInService) authenticateUser(username, password string) bool {
	return username == password
}

func (ss *SignInService) generateToken(request model.SignInRequest) (string, error) {
	token := base64.StdEncoding.EncodeToString([]byte(request.Username + request.Password))
	return token, nil
}
