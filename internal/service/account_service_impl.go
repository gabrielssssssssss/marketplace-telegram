package service

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *accountServiceImpl) Register(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.repository.Register(ctx, b, update)
}
