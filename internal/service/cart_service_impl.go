package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *cartServiceImpl) RegisterCart(cart *entity.Cart) (*model.Cart, error) {
	return s.repository.InsertCart(cart)
}

func (s *cartServiceImpl) FindCartByID(cart *entity.Cart) (*model.Cart, error) {
	return s.repository.SelectCartByID(cart)
}

func (s *cartServiceImpl) FindCartsByUserID(cart *entity.Cart) (*[]model.Cart, error) {
	return s.repository.SelectCartsByUserID(cart)
}

func (s *cartServiceImpl) ModifyCartByID(cart *entity.Cart) (*model.Cart, error) {
	return s.repository.UpdateCartByID(cart)
}

func (s *cartServiceImpl) RemoveCartByID(cart *entity.Cart) (bool, error) {
	return s.repository.DeleteCartByID(cart)
}
