package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type AccountService interface {
	RegisterUser(user *entity.Users) (*model.Users, error)
	FindUserByID(userID int64) (*model.Users, error)
	FindUserByRecoveryKey(recoveryKey string) (*model.Users, error)
	ModifyUserByID(user *entity.Users) (*model.Users, error)
	RemoveUserByID(UserID int64) (bool, error)
}

type accountServiceImpl struct {
	repository repository.UserRepository
}

func NewAccountService(userRepository repository.UserRepository) AccountService {
	return &accountServiceImpl{
		repository: userRepository,
	}
}
