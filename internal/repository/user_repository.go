package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type UserRepository interface {
	InsertUser(user *entity.User) (*model.User, error)
	SelectUserByID(user *entity.User) (*model.User, error)
	SelectUserByRecoveryKey(user *entity.User) (*model.User, error)
	UpdateUserByID(user *entity.User) (*model.User, error)
	DeleteUserByID(user *entity.User) (bool, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: client,
	}
}
