package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type ProductService interface {
	RegisterProduct(product *entity.Product) (*model.Product, error)
	FindProductByID(product *entity.Product) (*model.Product, error)
	ModifyProductByID(product *entity.Product) (*model.Product, error)
	RemoveProductByID(product *entity.Product) (bool, error)
}

type productServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(productRepository repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		repository: productRepository,
	}
}
