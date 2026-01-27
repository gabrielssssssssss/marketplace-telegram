package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *accountServiceImpl) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	req := &entity.Users{UserId: update.Message.Chat.ID}
	resp, err := s.repository.GetUserByID(req)

	if err != nil || resp.UserId != update.Message.Chat.ID {
		req = &entity.Users{
			UserId:      update.Message.Chat.ID,
			Username:    update.Message.Chat.Username,
			Firstname:   update.Message.Chat.FirstName,
			Lastname:    update.Message.Chat.LastName,
			Balance:     0,
			RecoveryKey: helper.RandomStringSecure(24),
			UpdatedAt:   time.Now(),
		}
		resp, _ = s.repository.CreateUser(req)
		message := fmt.Sprintf(messages.MessageRecoveryKey, resp.RecoveryKey)
		b.SendMessage(ctx, &bot.SendMessageParams{ChatID: update.Message.Chat.ID, Text: message})
	}

	b.SendMessage(ctx, &bot.SendMessageParams{ChatID: update.Message.Chat.ID, Text: "yo"})
}
