package model

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BuyBookRequest struct {
	BookId   string `json:"id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type OrderResponse struct{
	OrderId string `json:"orderId"`
}