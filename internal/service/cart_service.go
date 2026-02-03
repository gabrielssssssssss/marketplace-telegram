package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type CartService interface {
	RegisterCart(cart *entity.Cart) (*model.Cart, error)
	GetCartByID(cart *entity.Cart) (*model.Cart, error)
	GetCartsByUserID(cart *entity.Cart) (*[]model.Cart, error)
	ModifyCartByID(cart *entity.Cart) (*model.Cart, error)
	RemoveCartByID(cart *entity.Cart) (bool, error)
}

type cartServiceImpl struct {
	repository repository.CartRepository
}

func NewCartService(cartRepository repository.CartRepository) CartService {
	return &cartServiceImpl{
		repository: cartRepository,
	}
}
