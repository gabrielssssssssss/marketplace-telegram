package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type PaymentService interface {
	Register(payment *entity.Payment) (*model.Payment, error)
	GetPayment(payment *entity.Payment) (*model.Payment, error)
	Modify(payment *entity.Payment) (*model.Payment, error)
}

type paymentServiceImpl struct {
	repository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		repository: paymentRepository,
	}
}
