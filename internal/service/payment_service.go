package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type PaymentService interface {
}

type paymentServiceImpl struct {
	repository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		repository: paymentRepository,
	}
}
