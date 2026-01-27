package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

type AccountRepository interface {
	CreateUser(users *entity.Users) (*model.Users, error)
	GetUserByID(users *entity.Users) (*model.Users, error)
	DeleteUser(users *entity.Users) (bool, error)
}

type accountRepositoryImpl struct {
	db *sql.DB
}

func NewAccountRepository(client *sql.DB) AccountRepository {
	return &accountRepositoryImpl{
		db: client,
	}
}
