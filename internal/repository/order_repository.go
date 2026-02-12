package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type OrderRepository interface {
	Create(orders *entity.Order) (*model.Order, error)
	Order(orders *entity.Order) (*model.Order, error)
	UserOrders(orders *entity.Order) (*[]model.Order, error)
	Delete(orders *entity.Order) (bool, error)
}

type orderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepository(client *sql.DB) OrderRepository {
	return &orderRepositoryImpl{
		db: client,
	}
}
