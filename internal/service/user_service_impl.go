package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *userServiceImpl) RegisterUser(user *entity.User) (*model.User, error) {
	return s.repository.CreateUser(user)
}

func (s *userServiceImpl) FindUserByID(user *entity.User) (*model.User, error) {
	return s.repository.GetUserByID(user)
}

func (s *userServiceImpl) FindUserByRecoveryKey(user *entity.User) (*model.User, error) {
	return s.repository.GetUserByRecoveryKey(user)
}

func (s *userServiceImpl) ModifyUserByID(user *entity.User) (*model.User, error) {
	return s.repository.UpdateUserByID(user)
}

func (s *userServiceImpl) RemoveUserByID(user *entity.User) (bool, error) {
	return s.repository.DeleteUser(user)
}
