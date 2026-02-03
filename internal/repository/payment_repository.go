package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type PaymentRepository interface {
	InsertPayment(payment *entity.Payment) (*model.Payment, error)
	SelectPaymentByID(payment *entity.Payment) (*model.Payment, error)
	SelectPaymentsByUserID(payment *entity.Payment) (*[]model.Payment, error)
	UpdatePaymentByID(payment *entity.Payment) (*model.Payment, error)
	DeletePaymentByID(payment *entity.Payment) (bool, error)
}

type paymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepository(client *sql.DB) PaymentRepository {
	return &paymentRepositoryImpl{
		db: client,
	}
}
