package service

import (
	"bookstore-api/model"
	"bookstore-api/repository"
	"context"
)


type orderService struct{
	ordersRepo repository.OrdersRepository
}


type OrderService interface{
	GetAllOrders(ctx context.Context) (model.Orders, error)
}

func NewOrderService(ordersRepo repository.OrdersRepository)OrderService{
	return &orderService{
		ordersRepo: ordersRepo,
	}
}


func(o *orderService)GetAllOrders(ctx context.Context) (model.Orders, error) {

	return model.Orders{}, nil
}