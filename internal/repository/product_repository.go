package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type ProductRepository interface {
	Create(product *entity.Product) (*model.Product, error)
	Public(product *entity.Product) (*model.Product, error)
	Private(product *entity.Product) (*model.Product, error)
	List() (*[]model.Product, error)
	Update(product *entity.Product) (*model.Product, error)
	Delete(product *entity.Product) (bool, error)
}

type productRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(client *sql.DB) ProductRepository {
	return &productRepositoryImpl{
		db: client,
	}
}
