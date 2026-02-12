package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *cartServiceImpl) Register(cart *entity.Cart) (*model.Cart, error) {
	return s.repository.Create(cart)
}

func (s *cartServiceImpl) GetCart(cart *entity.Cart) (*model.Cart, error) {
	return s.repository.Cart(cart)
}

func (s *cartServiceImpl) GetCarts(cart *entity.Cart) (*[]model.Cart, error) {
	return s.repository.Carts(cart)
}

func (s *cartServiceImpl) Modify(cart *entity.Cart) (*model.Cart, error) {
	return s.repository.Update(cart)
}

func (s *cartServiceImpl) Remove(cart *entity.Cart) (bool, error) {
	return s.repository.Delete(cart)
}
