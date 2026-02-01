package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type AccountService interface {
	RegisterUser(user *entity.Users) (*model.User, error)
	FindUserByID(userID int64) (*model.User, error)
	FindUserByRecoveryKey(recoveryKey string) (*model.User, error)
	ModifyUserByID(user *entity.Users) (*model.User, error)
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
