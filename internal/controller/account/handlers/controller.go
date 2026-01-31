package account

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(AccountService *service.AccountService) AccountController {
	return AccountController{AccountService: *AccountService}
}

func (controller *AccountController) HandlerStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := controller.AccountService.StartCommand(ctx, b, update)

	if err != nil {
		log.Error().
			Err(err).
			Str("component", "AccountController.HandlerStart").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process start command")
		return
	}

	log.Info().Msg("start command processed successfully")
}

func (controller *AccountController) HandlerAccount(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := controller.AccountService.AccountCallback(ctx, b, update)

	if err != nil {
		log.Error().
			Err(err).
			Str("component", "AccountController.HandlerAccount").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process start callback")
		return
	}

	log.Info().Msg("account processed successfully")
}
