package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type OrderRepository interface {
	CreateOrder(orders *entity.Orders) (*model.Orders, error)
	GetOrderByID(orders *entity.Orders) (*model.Orders, error)
	GetOrdersByUserID(orders *entity.Orders) (*[]model.Orders, error)
	DeleteOrderByID(orders *entity.Orders) (bool, error)
}

type orderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepository(client *sql.DB) OrderRepository {
	return &orderRepositoryImpl{
		db: client,
	}
}
