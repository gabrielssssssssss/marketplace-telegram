package account

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
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
	user, err := handler.AccountService.FindUserByID(update.Message.From.ID)
	if err != nil {
		newUser := entity.Users{
			UserId:    update.Message.From.ID,
			Username:  update.Message.From.Username,
			Firstname: update.Message.From.FirstName,
			Lastname:  update.Message.From.LastName,
			Role:      "user",
		}

		ownerID, err := strconv.ParseInt(os.Getenv("OWNER_ID"), 10, 64)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "strconv.ParseInt").
				Int64("user_id", update.Message.From.ID).
				Msg("Failed to process int64 parser")
			return
		}

		if update.Message.From.ID == ownerID {
			newUser.Role = "admin"
		}

		user, err = handler.AccountService.RegisterUser(&newUser)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "AccountService.RegisterUser").
				Int64("user_id", update.Message.From.ID).
				Msg("Failed to process register user")
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			ParseMode: "HTML",
			Text:      fmt.Sprintf(messages.MessageRecoveryKey, user.RecoveryKey),
		})

		log.Info().Msg("register user processed successfully")
	}

	newSession, err := helper.NewJwtToken(user.UserId, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "helper.NewJwtToken").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process jwt generation")
		return
	}

	webAppUrl := os.Getenv("TELEGRAM_WEB_APP") + "?token=" + newSession
	fmt.Println(webAppUrl)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		ParseMode: "HTML",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "💳 Dêpot", CallbackData: "payment"},
					{Text: "🔑 Restauration", CallbackData: "restore"},
				}, {
					{Text: "🛍 Boutique", WebApp: &models.WebAppInfo{
						URL: webAppUrl,
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
			Str("component", "handler.HandlerAccount").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	user, err := handler.AccountService.FindUserByID(cb.Message.Message.Chat.ID)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "AccountController.HandlerAccount").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process start callback")
	}

	newSession, err := helper.NewJwtToken(user.UserId, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "helper.NewJwtToken").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process jwt generation")
		return
	}

	webAppUrl := os.Getenv("TELEGRAM_WEB_APP") + "?token=" + newSession

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    cb.Message.Message.Chat.ID,
		MessageID: cb.Message.Message.ID,
		ParseMode: "HTML",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "💳 Dêpot", CallbackData: "payment"},
					{Text: "🔑 Restauration", CallbackData: "restore"},
				}, {
					{Text: "🛍 Boutique", WebApp: &models.WebAppInfo{
						URL: webAppUrl,
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
