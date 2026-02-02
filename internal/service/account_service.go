package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type AccountService interface {
	RegisterUser(user *entity.User) (*model.User, error)
	FindUserByID(user *entity.User) (*model.User, error)
	FindUserByRecoveryKey(user *entity.User) (*model.User, error)
	ModifyUserByID(user *entity.User) (*model.User, error)
	RemoveUserByID(user *entity.User) (bool, error)
}

type accountServiceImpl struct {
	repository repository.UserRepository
}

func NewAccountService(userRepository repository.UserRepository) AccountService {
	return &accountServiceImpl{
		repository: userRepository,
	}
}
