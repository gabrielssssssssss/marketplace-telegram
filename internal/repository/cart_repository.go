package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type CartRepository interface {
	Create(cart *entity.Cart) (*model.Cart, error)
	Cart(cart *entity.Cart) (*model.Cart, error)
	Carts(cart *entity.Cart) (*[]model.Cart, error)
	Update(cart *entity.Cart) (*model.Cart, error)
	Delete(cart *entity.Cart) (bool, error)
}

type cartRepositoryImpl struct {
	db *sql.DB
}

func NewCartRepository(client *sql.DB) CartRepository {
	return &cartRepositoryImpl{
		db: client,
	}
}
