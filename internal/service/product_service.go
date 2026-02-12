package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type ProductService interface {
	Register(product *entity.Product) (*model.Product, error)
	GetPrivate(product *entity.Product) (*model.Product, error)
	GetPublic(product *entity.Product) (*model.Product, error)
	GetAll() (*[]model.Product, error)
	Modify(product *entity.Product) (*model.Product, error)
	Remove(product *entity.Product) (bool, error)
}

type productServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{
		repository: productRepository,
	}
}
