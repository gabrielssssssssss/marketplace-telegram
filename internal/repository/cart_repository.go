package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type CartRepository interface {
	CreateCart(cart *entity.Cart) (*model.Cart, error)
	GetCartByID(cart *entity.Cart) (*model.Cart, error)
	GetCartsByUserID(cart *entity.Cart) (*[]model.Cart, error)
	UpdateCartByID(cart *entity.Cart) (*model.Cart, error)
	DeleteCartByID(cart *entity.Cart) (bool, error)
}

type cartRepositoryImpl struct {
	db *sql.DB
}

func NewCartRepository(client *sql.DB) CartRepository {
	return &cartRepositoryImpl{
		db: client,
	}
}
