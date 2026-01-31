package account

import (
	"context"
	"fmt"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

type AccountHandler struct {
	AccountService service.AccountService
}

func NewAccountHandler(AccountService *service.AccountService) AccountHandler {
	return AccountHandler{AccountService: *AccountService}
}

func (handler *AccountHandler) HandlerStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := handler.AccountService.FindUser(update.Message.From.ID)
	if err == nil {
		newUser := entity.Users{
			UserId:    update.Message.From.ID,
			Username:  update.Message.From.Username,
			Firstname: update.Message.From.FirstName,
			Lastname:  update.Message.From.LastName,
		}

		user, err = handler.AccountService.RegisterUser(&newUser)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "PaymentController.HandlerPayment").
				Int64("user_id", update.Message.From.ID).
				Msg("Failed to process payment callback")
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			ParseMode: "HTML",
			Text:      fmt.Sprintf(messages.MessageRecoveryKey, user.RecoveryKey),
		})

		log.Info().Msg("register user processed successfully")
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
		Text: fmt.Sprintf(messages.MessageAccount,
			user.UserId,
			user.Username,
			user.Balance,
		),
	})

	log.Info().Msg("start command processed successfully")
}

func (handler *AccountHandler) HandlerAccount(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		log.Error().
			Err(fmt.Errorf("callback query or associated message is nil")).
			Str("component", "PaymentController.HandlerPayment").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	user, err := handler.AccountService.FindUser(cb.Message.Message.Chat.ID)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "AccountController.HandlerAccount").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process start callback")
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
		Text: fmt.Sprintf(messages.MessageAccount,
			user.UserId,
			user.Username,
			user.Balance,
		),
	})

	log.Info().Msg("account processed successfully")
}
