package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type OrderService interface {
	Register(order *entity.Order) (*model.Order, error)
	GetOrder(order *entity.Order) (*model.Order, error)
	GetUserOrders(order *entity.Order) (*[]model.Order, error)
	Remove(order *entity.Order) (bool, error)
}

type orderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderServiceImpl{
		repository: orderRepository,
	}
}
