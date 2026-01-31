package handlers

import (
	"context"
	"fmt"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

type RestoreHandler struct {
	AccountService service.AccountService
}

func NewRestoreHandler(AccountService *service.AccountService) RestoreHandler {
	return RestoreHandler{AccountService: *AccountService}
}

func (handler *RestoreHandler) HandlerRestore(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		log.Error().
			Err(fmt.Errorf("callback query or associated message is nil")).
			Str("component", "handler.HandlerRestore").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process restore callback")
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    cb.Message.Message.Chat.ID,
		MessageID: cb.Message.Message.ID,
		ParseMode: "HTML",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{{Text: "👈 Retour", CallbackData: "account"}},
			},
		},
		Text: messages.MessageRestore,
	})

	log.Info().Msg("account processed successfully")
}

func (handler *RestoreHandler) ListenerRestore(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := handler.AccountService.FindUserByRecoveryKey(update.Message.Text)
	if err != nil {
		return
	}

	if user.RecoveryKey != "" {
		message := fmt.Sprintf(messages.MessageRestoreConfirm,
			user.UserId,
			user.Username,
			user.Balance,
			user.CreatedAt,
			user.UpdatedAt,
		)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			ParseMode: "HTML",
			ReplyMarkup: &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{{Text: "✅ Transférer", CallbackData: "restore_transfer"}},
				},
			},
			Text: message,
		})
		log.Info().Msg("account processed successfully")
	}
}
