package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *productServiceImpl) RegisterProduct(product *entity.Product) (*model.Product, error) {
	return s.repository.InsertProduct(product)
}

func (s *productServiceImpl) GetProductPublic(product *entity.Product) (*model.Product, error) {
	return s.repository.SelectProductPublic(product)
}

func (s *productServiceImpl) GetProductPrivate(product *entity.Product) (*model.Product, error) {
	return s.repository.SelectProductPublic(product)
}

func (s *productServiceImpl) GetAllProducts() (*[]model.Product, error) {
	return s.repository.SelectAllProducts()
}

func (s *productServiceImpl) ModifyProductByID(product *entity.Product) (*model.Product, error) {
	return s.repository.UpdateProductByID(product)
}

func (s *productServiceImpl) RemoveProductByID(product *entity.Product) (bool, error) {
	return s.repository.DeleteProductByID(product)
}
