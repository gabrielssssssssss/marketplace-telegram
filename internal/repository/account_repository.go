package repository

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AccountRepository interface {
	Register(ctx context.Context, b *bot.Bot, update *models.Update)
}

type accountRepositoryImpl struct {
}

func NewAccountRepository() AccountRepository {
	return accountRepositoryImpl{}
}
