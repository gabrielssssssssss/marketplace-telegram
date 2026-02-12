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
	UserService service.UserService
}

func NewAccountHandler(UserService *service.UserService) AccountHandler {
	return AccountHandler{UserService: *UserService}
}

func (handler *AccountHandler) HandlerStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := handler.UserService.GetUser(&entity.User{UserId: update.Message.From.ID})
	if err != nil {
		newUser := entity.User{
			UserId:      update.Message.From.ID,
			Username:    update.Message.From.Username,
			Firstname:   update.Message.From.FirstName,
			Lastname:    update.Message.From.LastName,
			Role:        "user",
			Balance:     0.0,
			RecoveryKey: helper.RandomStringSecure(24),
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

		user, err = handler.UserService.Register(&newUser)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "handler.UserService.Register").
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

	handler.renderAccountMenu(ctx, b, nil, update.Message.From.ID)

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

	handler.renderAccountMenu(ctx, b, cb, cb.From.ID)

	log.Info().Msg("account processed successfully")
}

func (handler *AccountHandler) renderAccountMenu(ctx context.Context, b *bot.Bot, cb *models.CallbackQuery, chatID int64) {
	user, err := handler.UserService.GetUser(&entity.User{UserId: chatID})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "handler.UserService.GetUser").
			Int64("user_id", chatID).
			Msg("Failed to process start callback")
		return
	}

	newSession, err := helper.NewJwtToken(user.UserId, user.Role, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "helper.NewJwtToken").
			Int64("user_id", chatID).
			Msg("Failed to process jwt generation")
		return
	}
	fmt.Println(newSession)
	webAppUrl := os.Getenv("TELEGRAM_WEB_APP") + "?token=" + newSession
	markup := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "💳 Dêpot", CallbackData: "payment"},
				{Text: "🔑 Restauration", CallbackData: "restore"},
			}, {
				{Text: "🛍 Boutique", WebApp: &models.WebAppInfo{URL: webAppUrl}},
			},
		},
	}

	text := fmt.Sprintf(messages.MessageAccount, user.UserId, user.Username, user.Balance)

	if cb != nil {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      chatID,
			MessageID:   cb.Message.Message.ID,
			ParseMode:   "HTML",
			ReplyMarkup: markup,
			Text:        text,
		})
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      chatID,
			ParseMode:   "HTML",
			ReplyMarkup: markup,
			Text:        text,
		})
	}
}
