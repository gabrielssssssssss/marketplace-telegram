package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type UserRepository interface {
	CreateUser(user *entity.Users) (*model.Users, error)
	GetUserByID(user *entity.Users) (*model.Users, error)
	GetUserByRecoveryKey(user *entity.Users) (*model.Users, error)
	UpdateUserByID(user *entity.Users) (*model.Users, error)
	DeleteUser(user *entity.Users) (bool, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: client,
	}
}
