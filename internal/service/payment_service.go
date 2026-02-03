package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type PaymentService interface {
	RegisterPayment(payment *entity.Payment) (*model.Payment, error)
	GetPaymentByID(payment *entity.Payment) (*model.Payment, error)
	ModifyPaymentByID(payment *entity.Payment) (*model.Payment, error)
}

type paymentServiceImpl struct {
	repository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		repository: paymentRepository,
	}
}
