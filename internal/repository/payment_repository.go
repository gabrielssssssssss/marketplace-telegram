package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type PaymentRepository interface {
	CreateOrder(orders *entity.Orders) (*model.Orders, error)
	GetOrderByID(orders *entity.Orders) (*model.Orders, error)
	GetOrdersByUserID(orders *entity.Orders) (*[]model.Orders, error)
	DeleteOrderByID(orders *entity.Orders) (bool, error)
}

type paymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepository(client *sql.DB) PaymentRepository {
	return &paymentRepositoryImpl{
		db: client,
	}
}
