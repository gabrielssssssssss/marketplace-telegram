package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*model.User, error)
	GetUserByID(user *entity.User) (*model.User, error)
	GetUserByRecoveryKey(user *entity.User) (*model.User, error)
	UpdateUserByID(user *entity.User) (*model.User, error)
	DeleteUser(user *entity.User) (bool, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: client,
	}
}
