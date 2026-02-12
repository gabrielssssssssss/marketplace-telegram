package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type UserRepository interface {
	Create(user *entity.User) (*model.User, error)
	User(user *entity.User) (*model.User, error)
	UserByKey(user *entity.User) (*model.User, error)
	Update(user *entity.User) (*model.User, error)
	Delete(user *entity.User) (bool, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: client,
	}
}
