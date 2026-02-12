package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *paymentServiceImpl) Register(payment *entity.Payment) (*model.Payment, error) {
	return s.repository.Create(payment)
}

func (s *paymentServiceImpl) GetPayment(payment *entity.Payment) (*model.Payment, error) {
	return s.repository.Payment(payment)
}

func (s *paymentServiceImpl) Modify(payment *entity.Payment) (*model.Payment, error) {
	return s.repository.Update(payment)
}
