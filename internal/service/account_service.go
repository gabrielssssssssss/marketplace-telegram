package service

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AccountService interface {
	Start(ctx context.Context, b *bot.Bot, update *models.Update)
}

type accountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountServiceImpl{
		repository: accountRepository,
	}
}
