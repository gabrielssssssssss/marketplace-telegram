package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type OrderService interface {
	RegisterOrder(order *entity.Order) (*model.Order, error)
	FindOrderByID(order *entity.Order) (*model.Order, error)
	FindOrdersByUserID(order *entity.Order) (*[]model.Order, error)
	RemoveOrderByID(order *entity.Order) (bool, error)
}

type orderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderServiceImpl{
		repository: orderRepository,
	}
}
