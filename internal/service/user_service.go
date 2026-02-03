package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type UserService interface {
	RegisterUser(user *entity.User) (*model.User, error)
	GetUserByID(user *entity.User) (*model.User, error)
	GetUserByRecoveryKey(user *entity.User) (*model.User, error)
	ModifyUserByID(user *entity.User) (*model.User, error)
	RemoveUserByID(user *entity.User) (bool, error)
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: userRepository,
	}
}
