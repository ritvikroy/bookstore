package model

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BuyBookRequest struct {
	Id string `json:"id" binding:"required"`
}
