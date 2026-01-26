package account

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AccountController struct{}

func NewAccountController() AccountController {
	return AccountController{}
}

func (controller *AccountController) Register(ctx context.Context, b *bot.Bot, update *models.Update) {

}
