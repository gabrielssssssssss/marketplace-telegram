package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type AccountService interface {
	RegisterUser(*entity.Users) (*model.Users, error)
	FindUser(UserID int64) (*model.Users, error)
	UpdateUserBalance(user *entity.Users) (*model.Users, error)
}

type accountServiceImpl struct {
	repository repository.UserRepository
}

func NewAccountService(userRepository repository.UserRepository) AccountService {
	return &accountServiceImpl{
		repository: userRepository,
	}
}
