package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
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

	log.Info().Msg("restore processed successfully")
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

		data := fmt.Sprintf("restore_transfer_%s", user.RecoveryKey)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			ParseMode: "HTML",
			ReplyMarkup: &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{{Text: "✅ Transférer", CallbackData: data}},
				},
			},
			Text: message,
		})
		log.Info().Msg("restore listener processed successfully")
	}
}

func (handler *RestoreHandler) HandlerRestoreTransfer(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		log.Error().
			Err(fmt.Errorf("callback query or associated message is nil")).
			Str("component", "update.CallbackQuery").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process restore callback")
		return
	}

	recoveryKey := strings.Split(cb.Data, "_")[2]
	chatID := cb.Message.Message.Chat.ID

	userSender, err := handler.AccountService.FindUserByRecoveryKey(recoveryKey)
	if err != nil {
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: cb.ID,
			Text:            "🚫 La clé de récupération n'existe pas ou a déjà été utiliser.",
		})
		return
	}

	userRecipient, err := handler.AccountService.FindUserByID(chatID)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "handler.AccountService.FindUserByID").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process restore callback")
		return
	}

	userUpdate := entity.Users{
		UserId:    userRecipient.UserId,
		Balance:   userSender.Balance + userRecipient.Balance,
		UpdatedAt: time.Now(),
	}

	_, err = handler.AccountService.ModifyUserByID(&userUpdate)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "handler.AccountService.UpdateUserBalance").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process restore callback")
		return
	}

	_, err = handler.AccountService.RemoveUserByID(userSender.UserId)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "handler.AccountService.RemoveUserByID").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process restore callback")
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
		Text:            "✅ Compte restaurer avec succès!",
	})

	log.Info().Msg("restore_transfer processed successfully")
}
