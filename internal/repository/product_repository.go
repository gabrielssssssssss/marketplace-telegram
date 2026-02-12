package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type ProductRepository interface {
	InsertProduct(product *entity.Product) (*model.Product, error)
	SelectProductPublic(product *entity.Product) (*model.Product, error)
	SelectProductPrivate(product *entity.Product) (*model.Product, error)
	SelectAllProducts() (*[]model.Product, error)
	UpdateProductByID(product *entity.Product) (*model.Product, error)
	DeleteProductByID(product *entity.Product) (bool, error)
}

type productRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(client *sql.DB) ProductRepository {
	return &productRepositoryImpl{
		db: client,
	}
}
