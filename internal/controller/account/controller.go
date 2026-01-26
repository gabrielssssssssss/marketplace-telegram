package account

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(AccountService *service.AccountService) AccountController {
	return AccountController{AccountService: *AccountService}
}

func (controller *AccountController) Register(ctx context.Context, b *bot.Bot, update *models.Update) {
	controller.AccountService.Register(ctx, b, update)
}
