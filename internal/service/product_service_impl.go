package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *productServiceImpl) Register(product *entity.Product) (*model.Product, error) {
	return s.repository.Create(product)
}

func (s *productServiceImpl) GetPublic(product *entity.Product) (*model.Product, error) {
	return s.repository.Public(product)
}

func (s *productServiceImpl) GetPrivate(product *entity.Product) (*model.Product, error) {
	return s.repository.Private(product)
}

func (s *productServiceImpl) GetAll() (*[]model.Product, error) {
	return s.repository.List()
}

func (s *productServiceImpl) Modify(product *entity.Product) (*model.Product, error) {
	return s.repository.Update(product)
}

func (s *productServiceImpl) Remove(product *entity.Product) (bool, error) {
	return s.repository.Delete(product)
}
