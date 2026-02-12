package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type PaymentRepository interface {
	Create(payment *entity.Payment) (*model.Payment, error)
	Payment(payment *entity.Payment) (*model.Payment, error)
	UserPayment(payment *entity.Payment) (*[]model.Payment, error)
	Update(payment *entity.Payment) (*model.Payment, error)
	Delete(payment *entity.Payment) (bool, error)
}

type paymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepository(client *sql.DB) PaymentRepository {
	return &paymentRepositoryImpl{
		db: client,
	}
}
