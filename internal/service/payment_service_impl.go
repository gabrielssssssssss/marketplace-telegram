package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *paymentServiceImpl) CreatePayment(payment *entity.Payment) (*model.Payment, error) {
	return s.repository.InsertPayment(payment)
}

func (s *paymentServiceImpl) FindPayment(payment *entity.Payment) (*model.Payment, error) {
	return s.repository.SelectPaymentByID(payment)
}

func (s *paymentServiceImpl) ConfirmPayment(payment *entity.Payment) (*model.Payment, error) {
	return s.repository.UpdatePaymentByID(payment)
}
