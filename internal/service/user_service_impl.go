package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *userServiceImpl) Register(user *entity.User) (*model.User, error) {
	return s.repository.Create(user)
}

func (s *userServiceImpl) GetUser(user *entity.User) (*model.User, error) {
	return s.repository.User(user)
}

func (s *userServiceImpl) GetUserByKey(user *entity.User) (*model.User, error) {
	return s.repository.UserByKey(user)
}

func (s *userServiceImpl) Modify(user *entity.User) (*model.User, error) {
	return s.repository.Update(user)
}

func (s *userServiceImpl) Remove(user *entity.User) (bool, error) {
	return s.repository.Delete(user)
}
