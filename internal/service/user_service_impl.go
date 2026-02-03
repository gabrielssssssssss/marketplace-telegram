package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *userServiceImpl) RegisterUser(user *entity.User) (*model.User, error) {
	return s.repository.InsertUser(user)
}

func (s *userServiceImpl) FindUserByID(user *entity.User) (*model.User, error) {
	return s.repository.SelectUserByID(user)
}

func (s *userServiceImpl) FindUserByRecoveryKey(user *entity.User) (*model.User, error) {
	return s.repository.SelectUserByRecoveryKey(user)
}

func (s *userServiceImpl) ModifyUserByID(user *entity.User) (*model.User, error) {
	return s.repository.UpdateUserByID(user)
}

func (s *userServiceImpl) RemoveUserByID(user *entity.User) (bool, error) {
	return s.repository.DeleteUserByID(user)
}
