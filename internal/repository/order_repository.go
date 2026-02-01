package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type OrderRepository interface {
	CreateOrder(orders *entity.Order) (*model.Order, error)
	GetOrderByID(orders *entity.Order) (*model.Order, error)
	GetOrdersByUserID(orders *entity.Order) (*[]model.Order, error)
	DeleteOrderByID(orders *entity.Order) (bool, error)
}

type orderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepository(client *sql.DB) OrderRepository {
	return &orderRepositoryImpl{
		db: client,
	}
}
