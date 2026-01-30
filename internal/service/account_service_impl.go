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

func (s *accountServiceImpl) StartCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	users := &entity.Users{
		UserId: update.Message.Chat.ID,
	}

	resp, err := s.repository.GetUserByID(users)

	if err != nil {
		users = &entity.Users{
			UserId:      update.Message.Chat.ID,
			Username:    update.Message.Chat.Username,
			Firstname:   update.Message.Chat.FirstName,
			Lastname:    update.Message.Chat.LastName,
			Balance:     0.0,
			RecoveryKey: helper.RandomStringSecure(24),
			UpdatedAt:   time.Now(),
		}

		resp, _ = s.repository.CreateUser(users)

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
		Text: fmt.Sprintf(messages.MessageAccount, users.UserId, users.Username, users.Balance),
	})
}

func (s *accountServiceImpl) AccountCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	users := &entity.Users{
		UserId: cb.Message.Message.Chat.ID,
	}

	resp, err := s.repository.GetUserByID(users)
	if err != nil {
		return
	}

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    cb.Message.Message.Chat.ID,
		MessageID: cb.Message.Message.ID,
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
		Text: fmt.Sprintf(messages.MessageAccount, resp.UserId, resp.Username, resp.Balance),
	})
}
