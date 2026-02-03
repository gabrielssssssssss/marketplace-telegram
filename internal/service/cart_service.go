package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type CartService interface {
}

type cartServiceImpl struct {
	repository repository.CartRepository
}

func NewCartService(cartRepository repository.CartRepository) CartService {
	return &cartServiceImpl{
		repository: cartRepository,
	}
}
