package repository

import (
	"bookstore-api/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type ordersRepository struct {
	db *sql.DB
}

type OrdersRepository interface {
	CreateOrder(ctx context.Context, bookId, orderValue string, orderQty, count int) (orderID string, err error)
	GetAllOrders(ctx context.Context) (model.Orders, error)
}

func NewOrderRepository(db *sql.DB) OrdersRepository {
	return ordersRepository{
		db: db,
	}
}

const (
	CreateOrder = `insert into orders (userId, bookId, ordervalue, quantity) values($1, $2, $3, $4)`
)

func (b ordersRepository) CreateOrder(ctx context.Context, bookId, orderValue string, orderQty, updatedCount int) (orderID string, err error) {
	fail := func(err error) (string, error) {
		return "", errors.New(err.Error())
	}

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, UpdateInventory, updatedCount, bookId)
	if err != nil {
		return fail(err)
	}

	userId := 1
	result, err := tx.ExecContext(ctx, CreateOrder, userId, bookId, orderValue, orderQty)
	if err != nil {
		return fail(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fail(err)
	}
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	orderID = fmt.Sprintf("%v", id)

	return orderID, nil

}

func(o ordersRepository)GetAllOrders(ctx context.Context) (model.Orders, error){

	return model.Orders{}, nil
}
