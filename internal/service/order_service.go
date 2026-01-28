package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type OrderService interface {
}

type orderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderServiceImpl{
		repository: orderRepository,
	}
}
