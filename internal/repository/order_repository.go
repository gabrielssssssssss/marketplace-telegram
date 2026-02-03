package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type OrderRepository interface {
	InsertOrder(orders *entity.Order) (*model.Order, error)
	SelectOrderByID(orders *entity.Order) (*model.Order, error)
	SelectOrdersByUserID(orders *entity.Order) (*[]model.Order, error)
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
