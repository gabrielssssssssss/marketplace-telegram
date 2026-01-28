package payment

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type PaymentController struct {
}

func NewPaymentController() PaymentController {
	return PaymentController{}
}

func (controller *PaymentController) PaymentMenu(ctx context.Context, b *bot.Bot, update *models.Update) {

}
