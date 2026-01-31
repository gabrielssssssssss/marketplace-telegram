package payment

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type OrderController struct {
}

func NewOrderController() OrderController {
	return OrderController{}
}

func (controller *OrderController) OrderMenu(ctx context.Context, b *bot.Bot, update *models.Update) {

}
