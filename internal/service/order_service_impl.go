package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *orderServiceImpl) Register(order *entity.Order) (*model.Order, error) {
	return s.repository.Create(order)
}

func (s *orderServiceImpl) GetOrder(order *entity.Order) (*model.Order, error) {
	return s.repository.Order(order)
}

func (s *orderServiceImpl) GetUserOrders(order *entity.Order) (*[]model.Order, error) {
	return s.repository.UserOrders(order)
}

// func (s *orderServiceImpl) ModifyCartByID(cart *entity.Cart) (*model.Cart, error) {
// 	return s.repository.UpdateCartByID(cart)
// }

func (s *orderServiceImpl) Remove(order *entity.Order) (bool, error) {
	return s.repository.Delete(order)
}
