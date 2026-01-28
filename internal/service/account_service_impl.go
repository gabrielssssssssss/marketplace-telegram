package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *accountServiceImpl) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	req := &entity.Users{
		UserId: update.Message.Chat.ID,
	}

	resp, err := s.repository.GetUserByID(req)

	if err != nil {
		req = &entity.Users{
			UserId:      update.Message.Chat.ID,
			Username:    update.Message.Chat.Username,
			Firstname:   update.Message.Chat.FirstName,
			Lastname:    update.Message.Chat.LastName,
			Balance:     0.0,
			RecoveryKey: helper.RandomStringSecure(24),
			UpdatedAt:   time.Now(),
		}

		resp, _ = s.repository.CreateUser(req)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			ParseMode: "HTML",
			Text:      fmt.Sprintf(messages.MessageRecoveryKey, resp.RecoveryKey),
		})
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: "HTML",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "💳 Dêpot", CallbackData: "payment"},
					{Text: "🔑 Restauration", CallbackData: "recovery"},
				}, {
					{Text: "🛍 Boutique", WebApp: &models.WebAppInfo{
						URL: os.Getenv("TELEGRAM_WEB_APP"),
					}},
				},
			},
		},
		Text: fmt.Sprintf(messages.MessageMenu, req.UserId, req.Username, req.Balance),
	})
}
