package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
)

type UserService interface {
	Register(user *entity.User) (*model.User, error)
	GetUser(user *entity.User) (*model.User, error)
	GetUserByKey(user *entity.User) (*model.User, error)
	Modify(user *entity.User) (*model.User, error)
	Remove(user *entity.User) (bool, error)
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: userRepository,
	}
}
