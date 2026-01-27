package repository

import (
	"context"
	"database/sql"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AccountRepository interface {
	Start(ctx context.Context, b *bot.Bot, update *models.Update)
}

type accountRepositoryImpl struct {
	db *sql.DB
}

func NewAccountRepository(client *sql.DB) AccountRepository {
	return &accountRepositoryImpl{
		db: client,
	}
}
