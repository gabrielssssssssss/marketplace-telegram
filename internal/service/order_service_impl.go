package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *orderServiceImpl) RegisterOrder(order *entity.Order) (*model.Order, error) {
	return s.repository.InsertOrder(order)
}

func (s *orderServiceImpl) GetOrderByID(order *entity.Order) (*model.Order, error) {
	return s.repository.SelectOrderByID(order)
}

func (s *orderServiceImpl) GetOrdersByUserID(order *entity.Order) (*[]model.Order, error) {
	return s.repository.SelectOrdersByUserID(order)
}

// func (s *orderServiceImpl) ModifyCartByID(cart *entity.Cart) (*model.Cart, error) {
// 	return s.repository.UpdateCartByID(cart)
// }

func (s *orderServiceImpl) RemoveOrderByID(order *entity.Order) (bool, error) {
	return s.repository.DeleteOrderByID(order)
}
