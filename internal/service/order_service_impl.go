package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *orderServiceImpl) RegisterOrder(order *entity.Order) (*model.Order, error) {
	return s.repository.CreateOrder(order)
}

func (s *orderServiceImpl) FindOrderByID(order *entity.Order) (*model.Order, error) {
	return s.repository.GetOrderByID(order)
}

func (s *orderServiceImpl) FindOrdersByUserID(order *entity.Order) (*[]model.Order, error) {
	return s.repository.GetOrdersByUserID(order)
}

// func (s *orderServiceImpl) ModifyCartByID(cart *entity.Cart) (*model.Cart, error) {
// 	return s.repository.UpdateCartByID(cart)
// }

func (s *orderServiceImpl) RemoveOrderByID(order *entity.Order) (bool, error) {
	return s.repository.DeleteOrderByID(order)
}
