package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type CartService interface {
	Register(cart *entity.Cart) (*model.Cart, error)
	GetCart(cart *entity.Cart) (*model.Cart, error)
	GetCarts(cart *entity.Cart) (*[]model.Cart, error)
	Modify(cart *entity.Cart) (*model.Cart, error)
	Remove(cart *entity.Cart) (bool, error)
}

type cartServiceImpl struct {
	repository repository.CartRepository
}

func NewCartService(cartRepository repository.CartRepository) CartService {
	return &cartServiceImpl{
		repository: cartRepository,
	}
}
