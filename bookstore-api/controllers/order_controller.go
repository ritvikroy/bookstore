package controllers

import (
	"bookstore-api/model"
	"bookstore-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ordersController struct{
	orderService service.OrderService
}

type OrdersController interface{
	GetAllOrders(ctx *gin.Context)
}

func NewOrderController(orderService service.OrderService) OrdersController{
	return ordersController{
		orderService: orderService,
	}
}


func(o ordersController) GetAllOrders(ctx *gin.Context){



	ctx.JSON(http.StatusOK, model.Orders{})
}